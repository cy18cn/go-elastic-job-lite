package api

type ShardingContext struct {
	// job id
	JobId string
	// job name
	JobName string

	ShardingItem int32

	// sharding total count
	ShardingCount int32

	// Customized job parameter
	JobParameter interface{}
	// Customized ShardingItem parameter, key is shard
	ShardingItemParameter interface{}

	SendJobEvent                 bool
	JobEventSamplingCount        int
	CurrentJobEventSamplingCount int
}
