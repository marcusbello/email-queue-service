package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/marcusbello/email-queue-service/internal/email"
	"github.com/marcusbello/email-queue-service/internal/queue"
	"github.com/marcusbello/email-queue-service/internal/server"
	"github.com/marcusbello/email-queue-service/internal/worker"
)

const (
	DefaultHTTPAddr = "localhost:8080"
	DefaultQueueSize = 10
	DefaultNumberOfWorkers = 3
)

var numberOfWorkers int
var queueSize int
var httpAddr string

func init() {
	flag.IntVar(&numberOfWorkers, "workers", numberOfWorkers, "Number of worker goroutines")
	flag.IntVar(&queueSize, "queue_size", queueSize, "Queue size")
	flag.StringVar(&httpAddr, "haddr", DefaultHTTPAddr, "Set the HTTP address")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] \n", os.Args[0])
		flag.PrintDefaults()
	}
}


func main() {
	// Parse command line flags
	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		log.Printf("Using default values HTTP address: %s; workers: %d; Queue size: %d", httpAddr, numberOfWorkers, queueSize)
	} else {
		log.Printf("Using HTTP address: %s; workers: %d; Queue size: %d", httpAddr, numberOfWorkers, queueSize)
	}

	// Create an in-memory job queue
	var jobQueue queue.JobQueue = queue.NewInMemoryQueue(queueSize)

	// Create email sender
	emailSender := email.NewEmailSender()

	// start workers
	log.Printf("Starting %d workers...\n", numberOfWorkers)
	worker.StartWorkers(numberOfWorkers, jobQueue, emailSender)
	log.Println("Workers started successfully")
	
	// create server
	srv := server.NewServer(httpAddr, jobQueue)
	go srv.Start()

	// Handle shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
	jobQueue.Close()
	worker.WaitForWorkers()
	log.Println("Server stopped gracefully")
}