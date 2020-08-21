package api

// Elastic job Listener
type ElasticJobListener interface {
	// Run before job executed
	BeforeJobExecuted(ctx JobContext) error

	// Run after job executed
	AfterJobExecuted(ctx JobContext) error
}
