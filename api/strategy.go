package api

// Job sharding strategy
type JobShardingStrategy interface {
	// job sharding
	Sharding(instances []JobInstance, jobName string, shardingCount int32) map[JobInstance][]int32
}
