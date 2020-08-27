package config

type JobType int

const (
	SIMPLE   JobType = 1
	DATAFLOW JobType = 2
	SCRIPT   JobType = 3
)

type JobTypeConfiguration interface {
	// Get jobReflectType reflect type
	GetJobReflectType() string
	// Get jobReflectType type
	GetJobType() JobType
	// Get jobReflectType core configuration
	GetCoreConfig() *JobCoreConfiguration
}

type JobCoreConfiguration struct {
	jobName                string
	cron                   string
	shardingTotalCount     int32
	jobParameter           string
	shardingItemParameters string
	failover               bool
	misfire                bool
	description            string
}

func (config *JobCoreConfiguration) GetJobName() string {
	return config.jobName
}

func (config *JobCoreConfiguration) GetCron() string {
	return config.cron
}

func (config *JobCoreConfiguration) GetShardingTotalCount() int32 {
	return config.shardingTotalCount
}

func (config *JobCoreConfiguration) GetJobParameter() string {
	return config.jobParameter
}

func (config *JobCoreConfiguration) GetShardingItemParameters() string {
	return config.shardingItemParameters
}

func (config *JobCoreConfiguration) IsFailover() bool {
	return config.failover
}

func (config *JobCoreConfiguration) IsMisfire() bool {
	return config.misfire
}

func (config *JobCoreConfiguration) GetDescription() string {
	return config.description
}

type JobCoreConfigurationBuilder struct {
	JobName                string
	Cron                   string
	ShardingTotalCount     int32
	JobParameter           string
	ShardingItemParameters string
	Failover               bool
	// misfire在间隔时间短的任务中比较鸡肋。按照这个例子说，20s之后，新一次的作业就执行了，没必要再进行misfire。
	// misfire的正确用法是用于处理间隔时间长的作业，或者业务有局限的作业。举个例子：
	// 1天跑一次的报表作业，作业的业务逻辑中写死本次作业只抓取昨天的数据。那么一旦上次作业跑动超过一天，那么第二天本该运行的作业就失去了运行的机会，第二天的数据也会缺失。因此需要misfire。
	// 如果业务写的足够好，比如，无论哪天跑都能把未处理的数据分好组生成报表，misfire的存在意义就仅仅是触发时间缩短而已了。
	Misfire     bool
	Description string
}

func (builder *JobCoreConfigurationBuilder) build() JobCoreConfiguration {
	return JobCoreConfiguration{
		jobName:                builder.JobName,
		cron:                   builder.Cron,
		shardingTotalCount:     builder.ShardingTotalCount,
		jobParameter:           builder.JobParameter,
		shardingItemParameters: builder.ShardingItemParameters,
		failover:               builder.Failover,
		misfire:                builder.Misfire,
		description:            builder.Description,
	}
}

type SimpleJobConfiguration struct {
	jobCoreConfiguration JobCoreConfiguration
	jobReflectType       string
}

func NewSimpleJobConfiguration(coreConfig JobCoreConfiguration,
	jobReflectType string) SimpleJobConfiguration {
	return SimpleJobConfiguration{
		jobCoreConfiguration: coreConfig,
		jobReflectType:       jobReflectType,
	}
}

// jobReflectType reflect type
func (simple *SimpleJobConfiguration) GetJobReflectType() string {
	return simple.jobReflectType
}

func (simple *SimpleJobConfiguration) GetJobType() JobType {
	return SIMPLE
}

func (simple *SimpleJobConfiguration) GetCoreConfig() *JobCoreConfiguration {
	return &simple.jobCoreConfiguration
}

type DataflowJobConfiguration struct {
	jobCoreConfiguration JobCoreConfiguration
	jobReflectType       string
	streamProcessing     bool
}

func NewDataflowJobConfiguration(coreConfig JobCoreConfiguration,
	jobReflectType string) DataflowJobConfiguration {
	return DataflowJobConfiguration{
		jobCoreConfiguration: coreConfig,
		jobReflectType:       jobReflectType,
	}
}

func (dataflow *DataflowJobConfiguration) GetJobReflectType() string {
	return dataflow.jobReflectType
}

func (dataflow *DataflowJobConfiguration) GetJobType() JobType {
	return DATAFLOW
}

func (dataflow *DataflowJobConfiguration) GetCoreConfig() *JobCoreConfiguration {
	return &dataflow.jobCoreConfiguration
}

func (dataflow *DataflowJobConfiguration) StreamProcessing() bool {
	return dataflow.streamProcessing
}

type ScriptJobConfiguration struct {
	jobCoreConfiguration JobCoreConfiguration
	scriptCommandLine    string
}

func NewScriptJobConfiguration(coreConfig JobCoreConfiguration,
	scriptCommandLine string) ScriptJobConfiguration {
	return ScriptJobConfiguration{
		jobCoreConfiguration: coreConfig,
		scriptCommandLine:    scriptCommandLine,
	}
}

func (script *ScriptJobConfiguration) GetJobReflectType() string {
	return "api.ScriptJob"
}

func (script *ScriptJobConfiguration) GetJobType() JobType {
	return SCRIPT
}

func (script *ScriptJobConfiguration) GetCoreConfig() *JobCoreConfiguration {
	return &script.jobCoreConfiguration
}

func (script *ScriptJobConfiguration) ScriptCommandLine() string {
	return script.scriptCommandLine
}
