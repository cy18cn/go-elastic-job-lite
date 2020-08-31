package api

// Elastic job Listener
type ElasticJobListener interface {
	// Run before job executed
	BeforeJobExecuted(ctx JobContext) error

	// Run after job executed
	AfterJobExecuted(ctx JobContext) error
}

type DefaultJobListener struct {
}

func (d *DefaultJobListener) BeforeJobExecuted(ctx JobContext) error {
	return nil
}

func (d *DefaultJobListener) AfterJobExecuted(ctx JobContext) error {
	return nil
}
