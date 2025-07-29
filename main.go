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

	"github.com/marcusbello/email-queue-service/internal/queue"
	"github.com/marcusbello/email-queue-service/internal/server"
)

const (
	DefaultHTTPAddr = "localhost:8080"
)

var numberOfWorkers int = 3
var queueSize int = 10
var httpAddr string

func init() {
	flag.Int("workers", numberOfWorkers, "Number of worker goroutines")
	flag.Int("queue_size", queueSize, "Queue size")
	flag.StringVar(&httpAddr, "haddr", DefaultHTTPAddr, "Set the HTTP address")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] \n", os.Args[0])
		flag.PrintDefaults()
	}
}


func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		log.Printf("Using default values HTTP address: %s; workers: %d; Queue size: %d ", httpAddr, numberOfWorkers, queueSize)
	} else {
		log.Printf("Using HTTP address: %s; workers: %d; Queue size: %d ", httpAddr, numberOfWorkers, queueSize)
	}
	// Create a job queue
	var jobQueue queue.JobQueue = queue.NewInMemoryQueue(queueSize)

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
	log.Println("Server stopped gracefully")
}