package main

import (
	"log"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		log.Printf("Worker %d completed job %d", id, job)
		// Simulating some work
		result := job * 2
		log.Printf("Worker %d completed job %d", id, job)
		results <- result
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Create workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			worker(workerID, jobs, results)
			wg.Done()
		}(w)
	}

	// Assign jobs to workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()

	// Collect results
	close(results)
	for r := range results {
		log.Printf("Result: %d", r)
	}
}
