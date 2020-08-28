package scheduler

import (
	"fmt"
	"go-elastic-job-lite/api"
	"go-elastic-job-lite/executor"
	"reflect"
)

type LiteJob struct {
	job       api.ElasticJob
	jobFacade JobFacade
	executor  executor.JobExecutor
}

func NewLiteJob(
	job api.ElasticJob,
	jobFacade JobFacade,
) (*LiteJob, error) {
	var exec executor.JobExecutor
	var err error

	switch job.(type) {
	case api.SimpleJob:
		exec = executor.NewSimpleJobExecutor(jobFacade, job.(api.SimpleJob))
	case api.DataFlowJob:
		exec = executor.NewDataflowJobExecutor(jobFacade, job.(api.DataFlowJob))
	case api.ScriptJob:
		exec = executor.NewScriptJobExecutor(jobFacade)
	default:
		err = fmt.Errorf("Not support job type: %s\n", reflect.TypeOf(job).String())
	}

	if err != nil {
		return nil, err
	}

	return &LiteJob{
		job:       job,
		jobFacade: jobFacade,
		executor:  exec,
	}, nil
}

func (liteJob *LiteJob) Run() {
	liteJob.executor.Execute()
}
