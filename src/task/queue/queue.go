package queue

import (
	"context"
	"errors"
)

var (
	// ErrCancel indicates the task was cancelled.
	ErrCancel = errors.New("queue: task cancelled")

	// ErrNotFound indicates the task was not found in the queue.
	ErrNotFound = errors.New("queue: task not found")
)

type MapType map[string]string

// Task defines a unit of work in the queue.
type Task struct {
	// ID identifies this task.
	ID string `json:"id,omitempty"`

	// Data is the actual data in the entry.
	Data []byte `json:"data"`
}

// InfoT provides runtime information.
type InfoT struct {
	Pending []*Task `json:"pending"`
	Running []*Task `json:"running"`
	Agents []string `json:"agents"`
	Stats   struct {
		Workers  int `json:"worker_count"`
		Pending  int `json:"pending_count"`
		Running  int `json:"running_count"`
		Complete int `json:"completed_count"`
	} `json:"stats"`
}


// Filter filters tasks in the queue. If the Filter returns false,
// the Task is skipped and not returned to the subscriber.
type Filter func(*Task) bool

// Queue defines a task queue for scheduling tasks among
// a pool of workers.
type Queue interface {
	// Push pushes an task to the tail of this queue.
	Push(c context.Context, task *Task) error

	// Poll retrieves and removes a task head of this queue.
	Poll(c context.Context, f Filter, agent_id string) (*Task, error)

	// Extend extends the deadline for a task.
	Extend(c context.Context, id string) error

	// Done signals the task is complete.
	Done(c context.Context, id string) error

	// Error signals the task is complete with errors.
	Error(c context.Context, id string, err error) error

	// Evict removes a pending task from the queue.
	Evict(c context.Context, id string) error

	// Wait waits until the task is complete.
	Wait(c context.Context, id string) error

	// Info returns internal queue information.
	Info(c context.Context) InfoT
	GetRuning(c context.Context, project string) int
	GetWorkers(c context.Context) []*Task
}
