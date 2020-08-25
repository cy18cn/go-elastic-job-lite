package scheduler

import (
	"go-elastic-job-lite/api"
	"go-elastic-job-lite/config"
	"go-elastic-job-lite/event"
)

type JobFacade interface {
	// Load job configuration
	LoadJobRootConfiguration(fromCache bool) config.JobRootConfiguration

	// Check job execution environment
	CheckJobExecutionEnvironment() error

	// Failover if necessary
	FailoverIfNecessary() error

	// Register job begin
	RegisterJobBegin(ctx api.JobContext)

	// Register job completed
	RegisterJobCompleted(ctx api.JobContext)

	// Get current job context
	GetJobContext() api.JobContext

	// Set job miss execution tag, if has running items
	MisfireIfHasRunningItems(shardingItems []int32) bool

	// Clear job miss execution tag
	ClearMisfire(shardingItems []int32)

	// Execute misfire?
	HasExecuteMisfired(shardingItems []int32) bool

	// Eligible Job For Running
	// if need to stop or re-sharding or not stream will stop running
	EligibleForJobRunning() bool

	// re-sharding
	NeedReSharding() bool

	// before job executed
	BeforeExecuted(ctx api.JobContext) error

	// after job executed
	AfterExecuted(ctx api.JobContext) error

	// post job execution event
	PostJobExecutionEvent(event event.JobExecutionEvent)

	// Post job status trace event
	PostJobStatusTraceEvent(jobId, message string, state event.State)
}

type LiteJobFacade struct {
	listeners []api.ElasticJobListener
}

func (jobFacade *LiteJobFacade) LoadJobRootConfiguration(fromCache bool) config.JobRootConfiguration {
	panic("implement me")
}

func (jobFacade *LiteJobFacade) CheckJobExecutionEnvironment() error {
	panic("implement me")
}

func (jobFacade *LiteJobFacade) FailoverIfNecessary() error {
	panic("implement me")
}

func (jobFacade *LiteJobFacade) RegisterJobBegin(ctx api.JobContext) {
	panic("implement me")
}

func (jobFacade *LiteJobFacade) RegisterJobCompleted(ctx api.JobContext) {
	panic("implement me")
}

func (jobFacade *LiteJobFacade) GetJobContext() api.JobContext {
	panic("implement me")
}

func (jobFacade *LiteJobFacade) MisfireIfHasRunningItems(shardingItems []int32) bool {
	panic("implement me")
}

func (jobFacade *LiteJobFacade) ClearMisfire(shardingItems []int32) {
	panic("implement me")
}

func (jobFacade *LiteJobFacade) HasExecuteMisfired(shardingItems []int32) bool {
	panic("implement me")
}

func (jobFacade *LiteJobFacade) EligibleForJobRunning() bool {
	panic("implement me")
}

func (jobFacade *LiteJobFacade) NeedReSharding() bool {
	panic("implement me")
}

func (jobFacade *LiteJobFacade) BeforeExecuted(ctx api.JobContext) error {
	if jobFacade.listeners == nil {
		return nil
	}

	var err error
	for _, listener := range jobFacade.listeners {
		err = listener.BeforeJobExecuted(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (jobFacade *LiteJobFacade) AfterExecuted(ctx api.JobContext) error {
	if jobFacade.listeners == nil {
		return nil
	}

	var err error
	for _, listener := range jobFacade.listeners {
		err = listener.AfterJobExecuted(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (jobFacade *LiteJobFacade) PostJobExecutionEvent(event event.JobExecutionEvent) {
	panic("implement me")
}

func (jobFacade *LiteJobFacade) PostJobStatusTraceEvent(jobId, message string, state event.State) {
	panic("implement me")
}
