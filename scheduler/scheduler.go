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

func (scheduler *JobScheduler) Start() error {
	typeConfig := scheduler.jobConfig.GetTypeConfig()
	coreConfig := typeConfig.GetCoreConfig()

	liteJob, err := NewLiteJob(scheduler.job, &LiteJobFacade{})
	if err != nil {
		return err
	}

	scheduler.cron.AddJob(scheduler.jobConfig.GetJobName(), liteJob)
	scheduler.cron.Start()

	return nil
}
