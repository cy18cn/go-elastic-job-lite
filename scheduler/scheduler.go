package scheduler

import (
	"go-elastic-job-lite/api"
	"go-elastic-job-lite/config"
	"go-elastic-job-lite/reg"

	"github.com/robfig/cron/v3"
)

type JobScheduler struct {
	jobConfig config.LiteJobConfiguration
	regCenter *reg.Registry
	job       *api.ElasticJob
	cron      *cron.Cron
}

func (scheduler *JobScheduler) Start() {
	scheduler.cron.AddJob(scheduler.jobConfig.GetJobName(), scheduler.)
	scheduler.cron.Start()
}
