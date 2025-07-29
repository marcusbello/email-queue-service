package email

import (
	"log"
	"time"
)

type EmailJob struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type EmailSender interface {
	Send(EmailJob)
}

type LoggerEmailSender struct{}

func NewEmailSender() EmailSender {
	return &LoggerEmailSender{}
}

func (s *LoggerEmailSender) Send(job EmailJob) {
	log.Printf("Sending email to %s: %s\n", job.To, job.Subject)
	time.Sleep(1 * time.Second) // simulate email send
}
