package scheduler

import (
	"go-elastic-job-lite/api"
	"go-elastic-job-lite/executor"
)

type LiteJob struct {
	job        api.ElasticJob
	jobContext api.JobContext
}

func (liteJob *LiteJob) Run() {
	exec, err := executor.NewExecutor(liteJob.job)
	if err != nil {

	}
	exec.Execute()
}
