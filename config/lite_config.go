package config

type JobRootConfiguration interface {
	GetTypeConfig() JobTypeConfiguration
}

type LiteJobConfiguration struct {
	jobConfig           JobTypeConfiguration
	monitorExecution    bool
	monitorPort         int
	jobShardingStrategy string
	disabled            bool
	overwrite           bool

	maxDiffSeconds           int
	reconcileIntervalMinutes int
}

func (config *LiteJobConfiguration) GetTypeConfig() JobTypeConfiguration {
	return config.jobConfig
}

func (config *LiteJobConfiguration) GetMonitorExecution() bool {
	return config.monitorExecution
}

func (config *LiteJobConfiguration) GetMonitorPort() int {
	return config.monitorPort
}

func (config *LiteJobConfiguration) GetJobShardingStrategy() string {
	return config.jobShardingStrategy
}

func (config *LiteJobConfiguration) GetDisabled() bool {
	return config.disabled
}

func (config *LiteJobConfiguration) GetOverride() bool {
	return config.overwrite
}

func (config *LiteJobConfiguration) GetMaxDiffSeconds() int {
	return config.maxDiffSeconds
}

func (config *LiteJobConfiguration) GetReconcileIntervalMinutes() int {
	return config.reconcileIntervalMinutes
}

func (config *LiteJobConfiguration) GetJobName() string {
	return config.jobConfig.GetCoreConfig().GetJobName()
}

func (config *LiteJobConfiguration) IsFailover() bool {
	return config.jobConfig.GetCoreConfig().IsFailover()
}

// LiteJobConfiguration builder
type LiteJobConfigurationBuilder struct {
	JobConfig           JobTypeConfiguration
	MonitorExecution    bool
	MonitorPort         int
	JobShardingStrategy string
	Disabled            bool
	Overwrite           bool

	MaxDiffSeconds           int
	ReconcileIntervalMinutes int
}

func (builder *LiteJobConfigurationBuilder) Build() LiteJobConfiguration {
	return LiteJobConfiguration{
		jobConfig:                builder.JobConfig,
		monitorExecution:         builder.MonitorExecution,
		monitorPort:              builder.MonitorPort,
		jobShardingStrategy:      builder.JobShardingStrategy,
		disabled:                 builder.Disabled,
		overwrite:                builder.Overwrite,
		maxDiffSeconds:           builder.MaxDiffSeconds,
		reconcileIntervalMinutes: builder.ReconcileIntervalMinutes,
	}
}
