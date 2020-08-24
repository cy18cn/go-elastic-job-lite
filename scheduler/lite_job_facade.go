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
	ExecuteMisfired(shardingItems []int32) bool

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
}
