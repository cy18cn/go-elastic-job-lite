package executor

import (
	"fmt"
	"go-elastic-job-lite/api"
	"go-elastic-job-lite/config"
)

type ElasticJobExecutor interface {
	Execute() error
}

type BaseExecutor struct {
	JobName      string
	JobConfig    config.JobConfiguration
	ErrorHandler api.JobErrorHandler
	Listener     api.ElasticJobListener
}

func NewExecutor(job api.ElasticJob) (ElasticJobExecutor, error) {
	var jobExecutor ElasticJobExecutor

	switch job.(type) {
	case api.SimpleJob:
	default:
		return nil, fmt.Errorf("")
	}

	return jobExecutor, nil
}

type SimpleJobExecutor struct {
}

func (exec *SimpleJobExecutor) Execute() error {

}
