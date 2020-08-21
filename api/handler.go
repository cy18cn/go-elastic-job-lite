package api

type JobErrorHandler interface {
	// Handle exception
	HandleException(jobName string, err error)
}
