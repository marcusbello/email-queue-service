package worker

import (
	"log"
	"sync"

	"github.com/marcusbello/email-queue-service/internal/email"
	"github.com/marcusbello/email-queue-service/internal/queue"
)

var wg sync.WaitGroup

func StartWorkers(n int, queue queue.JobQueue, sender email.EmailSender) {
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			log.Printf("Worker %d started\n", id)
			for job := range queue.Jobs() {
				log.Printf("Worker %d processing job\n", id)
				sender.Send(job)
			}
			log.Printf("Worker %d shutting down\n", id)
		}(i + 1)
	}
}

func WaitForWorkers() {
	wg.Wait()
}
