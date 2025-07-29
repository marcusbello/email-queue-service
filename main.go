package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

}