package scheduler

import (
	"go-elastic-job-lite/config"
	"testing"
)

func TestJobScheduler_Start(t *testing.T) {
	conf := &config.LiteJobConfigurationBuilder{}
	s := JobScheduler{
		jobConfig: conf.Build(),
	}
}
