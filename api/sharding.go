package api

type ShardingContext struct {
	// job id
	jobId string
	// job name
	jobName string

	shardingItem int32

	// sharding total count
	shardingCount int32

	// Customized job parameter
	jobParameter string
	// Customized ShardingItem parameter, key is shard
	shardingItemParameter string

	sendJobEvent                 bool
	jobEventSamplingCount        int
	currentJobEventSamplingCount int
}

func NewShardingContext(
	jobId,
	jobName,
	jobParameter,
	shardingItemParameter string,
	shardingItem,
	shardingCount int32) ShardingContext {
	return ShardingContext{
		jobId:                 jobId,
		jobName:               jobName,
		shardingItem:          shardingItem,
		shardingCount:         shardingCount,
		jobParameter:          jobParameter,
		shardingItemParameter: shardingItemParameter,
	}
}

func (ctx ShardingContext) JobId() string {
	return ctx.jobId
}

func (ctx ShardingContext) JobName() string {
	return ctx.jobName
}

func (ctx ShardingContext) ShardingItem() int32 {
	return ctx.shardingItem
}

func (ctx ShardingContext) ShardingCount() int32 {
	return ctx.shardingCount
}

func (ctx ShardingContext) JobParameter() string {
	return ctx.jobParameter
}

func (ctx ShardingContext) ShardingItemParameter() string {
	return ctx.shardingItemParameter
}
