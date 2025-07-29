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

	"github.com/marcusbello/email-queue-service/internal/server"
)

const (
	DefaultHTTPAddr = "localhost:8080"
)

var numberOfWorkers int = 3
var httpAddr string

func init() {
	flag.Int("workers", numberOfWorkers, "Number of worker goroutines")
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
		log.Printf("Using default HTTP address: %s with %d Workers ", httpAddr, numberOfWorkers)
	}
	log.Printf("Using HTTP address: %s with %d Workers ", httpAddr, numberOfWorkers)

	// create server
	srv := server.NewServer(httpAddr)
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