package schedule

import (
	"go-elastic-job-lite/api"
	"go-elastic-job-lite/config"
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

	// Set job miss execution tag
	MisfireIfRunning(shardingItems []int32)

	// Clear job miss execution tag
	ClearMisfire(shardingItems []int32)

	// Execute misfire?
	ExecuteMisfired(shardingItems []int32) bool

	// Eligible Job For Running
	// if need to stop or re-sharding or not stream will stop running
	EligibleForJobRunning() bool
}
