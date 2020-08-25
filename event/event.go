package event

import (
	"go-elastic-job-lite/api"
	"go-elastic-job-lite/util"
	"os"
	"time"
)

type ExecutionSource int8

// Trace event source
type Source int8

// Trace event state
type State int8

const (
	NORMAL_TRIGGER ExecutionSource = 1
	MISFIRE        ExecutionSource = 2
	FAILOVER       ExecutionSource = 3

	TASK_STAGING State = 1
	TASK_RUNNING
	TASK_FINISHED
	TASK_KILLED
	TASK_LOST
	TASK_FAILED
	TASK_ERROR
	TASK_DROPPED
	TASK_GONE
	TASK_GONE_BY_OPERATOR
	TASK_UNREACHABLE
	TASK_UNKNOWN

	CLOUD_SCHEDULER Source = 1
	CLOUD_EXECUTOR
	LITE_EXECUTOR
)

// Job event
type JobEvent interface {
	// get job name
	GetJobName() string
}

type JobExecutionEvent struct {
	id           string
	hostName     string
	ip           string
	jobId        string
	jobName      string
	source       ExecutionSource
	shardingItem int32
	startAt      time.Time
	completedAt  time.Time
	success      bool
	failureCause string
}

func NewJobExecutionEvent(
	jobId,
	jobName string,
	source ExecutionSource,
	shardingItem int32) (*JobExecutionEvent, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	ip, err := util.HostIP()
	if err != nil {
		return nil, err
	}

	return &JobExecutionEvent{
		id:           util.UUID(),
		hostName:     hostname,
		ip:           ip,
		jobId:        jobId,
		jobName:      jobName,
		source:       source,
		shardingItem: shardingItem,
		startAt:      time.Now(),
	}, nil
}

func (event *JobExecutionEvent) GetJobName() string {
	return event.jobName
}

func (event *JobExecutionEvent) ExecutionSuccess() *JobExecutionEvent {
	event.success = true
	event.completedAt = time.Now()
	return event
}

func (event *JobExecutionEvent) ExecutionFailure(cause string) *JobExecutionEvent {
	event.success = false
	event.failureCause = cause
	event.completedAt = time.Now()
	return event
}

func (event *JobExecutionEvent) FailureCause() string {
	return event.failureCause
}

type JobStatusTraceEvent struct {
	id            string
	originalJobId string
	jobId         string
	jobName       string
	slaveId       string
	source        Source
	executionType api.ExecutionType
	shardingItems string
	state         State
	createdAt     time.Time
	message       string
}

func NewTraceEvent(
	jobName,
	jobId,
	originalJobId,
	slaveId,
	shardingItems,
	message string,
	source Source,
	state State,
	executionType api.ExecutionType) *JobStatusTraceEvent {
	return &JobStatusTraceEvent{
		id:            util.UUID(),
		jobName:       jobName,
		jobId:         jobId,
		originalJobId: originalJobId,
		slaveId:       slaveId,
		shardingItems: shardingItems,
		message:       message,
		source:        source,
		state:         state,
		createdAt:     time.Now(),
		executionType: executionType,
	}
}

func (event *JobStatusTraceEvent) GetJobName() string {
	return event.jobName
}
