package api

type ExecutionType int8

const (
	// ready to execute
	READY ExecutionType = 1
	// failover task
	FAILOVER
)

type ElasticJob interface {
}

type JobInstance struct {
	JobInstanceId string
}

type SimpleJob interface {
	Execute(ctx ShardingContext) error
}

type DataFlowJob interface {
	FetchData(ctx ShardingContext) []interface{}
	ProcessData(ctx ShardingContext, data []interface{}) error
}

type ScriptJob interface {
}

type JobContext struct {
	// job id
	jobId string
	// job name
	jobName string
	// sharding total count
	shardingCount int32

	shardingItems []int32

	// Customized job parameter
	jobParameter string
	// Customized ShardingItem parameters, key is shard
	shardingItemParameters map[int32]string

	sendJobEvent                 bool
	jobEventSamplingCount        int
	currentJobEventSamplingCount int
}

func NewJobContext(
	jobId, jobName, jobParameter string,
	shardingCount int32,
	shardingItems []int32,
	shardingItemParameters map[int32]string,
	jobEventSamplingCount int,
) JobContext {
	return JobContext{
		jobId:                  jobId,
		jobName:                jobName,
		shardingCount:          shardingCount,
		shardingItems:          shardingItems,
		jobParameter:           jobParameter,
		shardingItemParameters: shardingItemParameters,
		jobEventSamplingCount:  jobEventSamplingCount,
	}
}

func (ctx *JobContext) SetShardingItems(shardingItems []int32) {
	ctx.shardingItems = shardingItems
}

func (ctx *JobContext) ShardingItems() []int32 {
	return ctx.shardingItems
}

func (ctx *JobContext) JobId() string {
	return ctx.jobId
}

func (ctx *JobContext) JobName() string {
	return ctx.jobName
}

func (ctx *JobContext) ShardingCount() int32 {
	return ctx.shardingCount
}

func (ctx *JobContext) JobParameter() string {
	return ctx.jobParameter
}

func (ctx *JobContext) ShardingItemParameters() map[int32]string {
	return ctx.shardingItemParameters
}

func (ctx *JobContext) SendJobEvent() bool {
	return ctx.sendJobEvent
}

func (ctx *JobContext) SetSendJobEvent(sendJobEvent bool) {
	ctx.sendJobEvent = sendJobEvent
}

func (ctx *JobContext) JobEventSamplingCount() int {
	return ctx.jobEventSamplingCount
}

func (ctx *JobContext) CurrentJobEventSamplingCount() int {
	return ctx.currentJobEventSamplingCount
}

func (ctx *JobContext) SetCurrentJobEventSamplingCount(currentJobEventSamplingCount int) {
	ctx.currentJobEventSamplingCount = currentJobEventSamplingCount
}
