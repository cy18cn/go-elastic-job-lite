package executor

import (
	"go-elastic-job-lite/api"
	"go-elastic-job-lite/config"
	"go-elastic-job-lite/event"
	"go-elastic-job-lite/scheduler"
	"os/exec"
	"sync"
)

type JobProcessor func(ctx api.ShardingContext) error

type ElasticJobExecutor struct {
	jobName      string
	jobFacade    scheduler.JobFacade
	jobConfig    config.JobRootConfiguration
	errorHandler api.JobErrorHandler
	processor    JobProcessor
}

func (exec *ElasticJobExecutor) Execute() {
	jobFacade := exec.jobFacade
	var err error
	// 1.check execution env
	//err = jobFacade.CheckJobExecutionEnvironment()
	//if err != nil {
	//	exec.errorHandler.HandleException(exec.jobName, err)
	//}
	//
	jobContext := jobFacade.GetJobContext()
	//if jobContext.SendJobEvent {
	//	jobFacade.PostJobStatusTraceEvent(
	//		jobContext.JobId, fmt.Sprintf("Job %s execute bigin", exec.jobName), event.TASK_STAGING)
	//}

	// 2. Add as misfire job if previous is still running
	//if jobFacade.MisfireIfHasRunningItems(jobContext.ShardingItems()) {
	//	if jobContext.SendJobEvent {
	//		jobFacade.PostJobStatusTraceEvent(
	//			jobContext.JobId,
	//			fmt.Sprintf("Previous job '%s'-shardingItems '%v' is still running,"+
	//				" misfired job will start after previous job completed.",
	//				exec.jobName,
	//				jobContext.ShardingItems()),
	//			event.TASK_FINISHED)
	//	}
	//	return
	//}

	// 3. Execute job
	err = jobFacade.BeforeExecuted(jobContext)
	if err != nil {
		exec.errorHandler.HandleException(jobContext.JobName, err)
	}
	exec.execute(jobContext, event.NORMAL_TRIGGER)

	// 4. Execute misfire job
	//shardingItems := jobContext.ShardingItems()
	//for jobFacade.HasExecuteMisfired(shardingItems) {
	//	jobFacade.ClearMisfire(shardingItems)
	//	exec.execute(jobContext, event.MISFIRE)
	//}
	//
	//jobFacade.FailoverIfNecessary()

	err = jobFacade.AfterExecuted(jobContext)
	if err != nil {
		exec.errorHandler.HandleException(jobContext.JobName, err)
	}
}

func (exec *ElasticJobExecutor) execute(ctx api.JobContext, source event.ExecutionSource) {
	//jobFacade := exec.jobFacade

	if ctx.ShardingItems() == nil {
		//if ctx.SendJobEvent {
		//	jobFacade.PostJobStatusTraceEvent(
		//		ctx.JobId,
		//		fmt.Sprintf("Sharding item for job '%s' is empty.", exec.jobName),
		//		event.TASK_FINISHED)
		//}
		return
	}

	// 1. register job begin
	//jobFacade.RegisterJobBegin(ctx)
	//jobId := ctx.JobId
	//if ctx.SendJobEvent {
	//	jobFacade.PostJobStatusTraceEvent(
	//		jobId,
	//		"",
	//		event.TASK_FINISHED)
	//}

	// 2. execute sharding job
	exec.process(ctx, source)

	// 3. register job completed
	//jobFacade.RegisterJobCompleted(ctx)
}

func (exec *ElasticJobExecutor) process(ctx api.JobContext, source event.ExecutionSource) {
	//jobFacade := exec.jobFacade
	shardingItems := ctx.ShardingItems()
	wg := sync.WaitGroup{}
	wg.Add(len(shardingItems))
	for _, item := range shardingItems {
		//startEvent, err := event.NewJobExecutionEvent(ctx.JobId, ctx.JobName, source, item)
		//if err != nil {
		//	exec.errorHandler.HandleException(ctx.JobName, err)
		//	return
		//}
		go func() {
			defer wg.Done()
			//if ctx.SendJobEvent {
			//	jobFacade.PostJobExecutionEvent(*startEvent)
			//}

			err := exec.processor(api.NewShardingContext(
				ctx.JobId,
				ctx.JobName,
				ctx.JobParameter,
				ctx.ShardingItemParameters[item],
				item,
				ctx.ShardingCount,
			))
			if err != nil {
				//completeEvent := startEvent.ExecutionFailure(err.Error())
				//jobFacade.PostJobExecutionEvent(*completeEvent)
				exec.errorHandler.HandleException(ctx.JobName, err)
				return
			}

			//completeEvent := startEvent.ExecutionSuccess()
			//if ctx.SendJobEvent {
			//	jobFacade.PostJobExecutionEvent(*completeEvent)
			//}
		}()
	}

	wg.Wait()
}

type SimpleJobExecutor struct {
	ElasticJobExecutor
}

func NewSimpleJobExecutor(
	jobFacade scheduler.JobFacade,
	job api.SimpleJob,
) *SimpleJobExecutor {
	jobConfig := jobFacade.LoadJobRootConfiguration(true)
	jobName := jobConfig.GetTypeConfig().GetCoreConfig().GetJobName()
	//errorHandler = jobConfig.GetTypeConfig().GetCoreConfig().

	return &SimpleJobExecutor{
		ElasticJobExecutor: ElasticJobExecutor{
			jobName:      jobName,
			jobFacade:    jobFacade,
			jobConfig:    jobConfig,
			errorHandler: nil,
			processor: func(ctx api.ShardingContext) error {
				return job.Execute(ctx)
			},
		},
	}
}

type DataflowJobExecutor struct {
	ElasticJobExecutor
}

func NewDataflowJobExecutor(
	jobFacade scheduler.JobFacade,
	job api.DataFlowJob,
) *SimpleJobExecutor {
	jobConfig := jobFacade.LoadJobRootConfiguration(true)
	jobName := jobConfig.GetTypeConfig().GetCoreConfig().GetJobName()
	//errorHandler = jobConfig.GetTypeConfig().GetCoreConfig().

	return &SimpleJobExecutor{
		ElasticJobExecutor: ElasticJobExecutor{
			jobName:      jobName,
			jobFacade:    jobFacade,
			jobConfig:    jobConfig,
			errorHandler: nil,
			processor: func(ctx api.ShardingContext) error {
				dataflowConfig := jobConfig.GetTypeConfig().(*config.DataflowJobConfiguration)
				var err error
				if dataflowConfig.StreamProcessing() {
					data := job.FetchData(ctx)
					for data != nil && len(data) > 0 {
						err = job.ProcessData(ctx, data)
						if err != nil {

						}
						data = job.FetchData(ctx)
					}
				} else {
					data := job.FetchData(ctx)
					if data != nil {
						err = job.ProcessData(ctx, data)
					}
				}

				return err
			},
		},
	}
}

type ScriptJobExecutor struct {
	ElasticJobExecutor
}

func NewScriptJobExecutor(
	jobFacade scheduler.JobFacade,
) *SimpleJobExecutor {
	jobConfig := jobFacade.LoadJobRootConfiguration(true)
	jobName := jobConfig.GetTypeConfig().GetCoreConfig().GetJobName()
	//errorHandler = jobConfig.GetTypeConfig().GetCoreConfig().

	return &SimpleJobExecutor{
		ElasticJobExecutor: ElasticJobExecutor{
			jobName:      jobName,
			jobFacade:    jobFacade,
			jobConfig:    jobConfig,
			errorHandler: nil,
			processor: func(ctx api.ShardingContext) error {
				commandLine := jobConfig.GetTypeConfig().(*config.ScriptJobConfiguration).ScriptCommandLine()
				cmd := exec.Command("/bin/bash", "-c", commandLine)
				_, err := cmd.Output()
				return err
			},
		},
	}
}
