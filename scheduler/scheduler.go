package scheduler

import (
	"go-elastic-job-lite/api"
	"go-elastic-job-lite/config"
	"go-elastic-job-lite/reg"
	"go-elastic-job-lite/sharding"

	"github.com/robfig/cron/v3"
)

type JobScheduler struct {
	jobConfig   config.LiteJobConfiguration
	regCenter   *reg.Registry
	job         *api.ElasticJob
	cron        *cron.Cron
	shardingSvc *sharding.Service
	listeners   []api.ElasticJobListener
}

func (scheduler *JobScheduler) Start() error {
	config := scheduler.jobConfig.GetTypeConfig().GetCoreConfig()
	ctx := api.NewJobContext(
		config.GetJobName(),
		config.GetJobParameter(),
		config.GetShardingTotalCount(),
		scheduler.shardingSvc.ShardingItems(),
		config.GetShardingItemParameters(),
		80,
	)
	liteJob, err := NewLiteJob(scheduler.job, NewJobFacade(ctx, scheduler.listeners))
	if err != nil {
		return err
	}

	scheduler.cron.AddJob(scheduler.jobConfig.GetJobName(), liteJob)
	scheduler.cron.Start()

	return nil
}
