package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
	"sync"
)

var (
	url          = kingpin.Arg("url", "URL to test").Required().String()
	verbose      = kingpin.Flag("verbose", "Verbose mode, shows requests").Short('v').Bool()
	reqPerWorker = kingpin.Flag("worker-requests", "number of requests per worker, use 0 for infinite").Default("0").Int()
	numWorkers   = kingpin.Flag("workers", "number of workers").Default("1000").Int()
)

func main() {

	kingpin.Parse()
	var wg sync.WaitGroup

	for i := 0; i < *numWorkers; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}
	wg.Wait()
}

func worker(wg *sync.WaitGroup, id int) {
	fmt.Println("Started worker ", id)

	defer wg.Done()

	for i := 0; i < *reqPerWorker || *reqPerWorker == 0; i++ {
		resp, _ := http.Get(*url)
		if *verbose {
			fmt.Println("[WORKER", id, "]:", resp)
		}
	}
}
