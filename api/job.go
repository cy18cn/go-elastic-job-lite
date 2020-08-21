package api

type ElasticJob interface {
}

type JobInstance struct {
	JobInstanceId string
}

type SimpleJob interface {
	Execute(ctx ShardingContext)
}

type DataFlowJob interface {
	FetchData(ctx ShardingContext) []interface{}
	ProcessData(ctx ShardingContext, data []interface{})
}

type ScriptJob interface {
}

type JobContext struct {
	// job id
	JobId string
	// job name
	JobName string
	// sharding total count
	ShardingCount int32

	// Customized job parameter
	JobParameter interface{}
	// Customized ShardingItem parameters, key is shard
	ShardingItemParameters map[int32]interface{}

	SendJobEvent                 bool
	JobEventSamplingCount        int
	CurrentJobEventSamplingCount int
}
