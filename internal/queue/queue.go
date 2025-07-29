package queue

import "github.com/marcusbello/email-queue-service/internal/email"

type JobQueue interface {
	Enqueue(job email.EmailJob) error
	Jobs() <-chan email.EmailJob
	Close()
}

