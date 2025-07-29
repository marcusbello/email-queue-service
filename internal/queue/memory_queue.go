package queue

import (
	"errors"

	"github.com/marcusbello/email-queue-service/internal/email"
)

type InMemoryQueue struct {
	jobs chan email.EmailJob
}

// NewInMemoryQueue creates a new in-memory job queue with the specified size.
// Using a buffered channel to simulate a queue with the specified size.
func NewInMemoryQueue(size int) *InMemoryQueue {
	return &InMemoryQueue{jobs: make(chan email.EmailJob, size)}
}

func (q *InMemoryQueue) Enqueue(job email.EmailJob) error {
	select {
	case q.jobs <- job:
		return nil
	default:
		return errors.New("queue full")
	}
}

func (q *InMemoryQueue) Jobs() <-chan email.EmailJob {
	return q.jobs
}

func (q *InMemoryQueue) Close() {
	close(q.jobs)
}
