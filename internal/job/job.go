type Status string

const (
	StatusQueued    Status = "QUEUED"
	StatusRunning   Status = "RUNNING"
	StatusCompleted Status = "COMPLETED"
	StatusFailed    Status = "FAILED"
)

type Job struct {
	ID       string
	TenantID string

	Status   Status
	Priority int

	Payload string
}
