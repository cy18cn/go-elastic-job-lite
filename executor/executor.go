package executor

import (
	"fmt"
	"go-elastic-job-lite/api"
	"go-elastic-job-lite/config"
	"go-elastic-job-lite/event"
	"go-elastic-job-lite/scheduler"
	"sync"
)

type JobProcessor func(ctx api.ShardingContext) error

type ElasticJobExecutor struct {
	jobName      string
	jobFacade    scheduler.JobFacade
	jobConfig    config.JobRootConfiguration
	errorHandler api.JobErrorHandler
	listener     api.ElasticJobListener
	processor    JobProcessor
}

func (exec *ElasticJobExecutor) Execute() error {
	// check execution env
	err := exec.jobFacade.CheckJobExecutionEnvironment()
	if err != nil {
		exec.errorHandler.HandleException(exec.jobName, err)
	}

	jobContext := exec.jobFacade.GetJobContext()
	if jobContext.SendJobEvent {
		exec.jobFacade.PostJobStatusTraceEvent(
			jobContext.JobId, fmt.Sprintf("Job %s execute bigin", exec.jobName), event.TASK_STAGING)
	}

	if exec.jobFacade.MisfireIfHasRunningItems(jobContext.ShardingItems()) {
		if jobContext.SendJobEvent {
			exec.jobFacade.PostJobStatusTraceEvent(
				jobContext.JobId,
				fmt.Sprintf("Previous job '%s'-shardingItems '%v' is still running,"+
					" misfired job will start after previous job completed.",
					exec.jobName,
					jobContext.ShardingItems()),
				event.TASK_FINISHED)
		}
		return nil
	}

	err = exec.jobFacade.BeforeExecuted(jobContext)
	if err != nil {
		exec.errorHandler.HandleException(jobContext.JobName, err)
	}
	exec.execute(jobContext, event.NORMAL_TRIGGER)
}

func (exec *ElasticJobExecutor) execute(ctx api.JobContext, source event.ExecutionSource) error {
	if ctx.ShardingItems() == nil {
		if ctx.SendJobEvent {
			exec.jobFacade.PostJobStatusTraceEvent(
				ctx.JobId,
				fmt.Sprintf("Sharding item for job '%s' is empty.", exec.jobName),
				event.TASK_FINISHED)
		}
		return nil
	}

	exec.jobFacade.RegisterJobBegin(ctx)
	jobId := ctx.JobId
	if ctx.SendJobEvent {
		exec.jobFacade.PostJobStatusTraceEvent(
			jobId,
			"",
			event.TASK_FINISHED)
	}

	shardingItems := ctx.ShardingItems()
	wg := sync.WaitGroup{}
	wg.Add(len(shardingItems))
	for _, item := range shardingItems {
		startEvent, err := event.NewJobExecutionEvent(ctx.JobId, ctx.JobName, source, item)
		if err != nil {
			return err
		}
		go func() {
			defer wg.Done()
			if ctx.SendJobEvent {
				exec.jobFacade.PostJobExecutionEvent(*startEvent)
			}

			err = exec.processor(api.ShardingContext{})
			if err != nil {
				completeEvent := startEvent.ExecutionFailure(err.Error())
				exec.jobFacade.PostJobExecutionEvent(*completeEvent)
				exec.errorHandler.HandleException(ctx.JobName, err)
				return
			}

			completeEvent := startEvent.ExecutionSuccess()
			if ctx.SendJobEvent {
				exec.jobFacade.PostJobExecutionEvent(*completeEvent)
			}
		}()
	}

	wg.Wait()

	return nil
}
