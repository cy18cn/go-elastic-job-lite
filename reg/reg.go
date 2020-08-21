package reg

import (
	"go-elastic-job-lite/api"
	"go-elastic-job-lite/config"
)

type Registry interface {
	Namespace(namespace string) error

	ApplyConfig(jobConfig config.JobConfiguration) error

	RegisterJob(job api.ElasticJob) error

	JobBegin(cxt api.JobContext) error

	JobCompleted(cxt api.JobContext) error
}
