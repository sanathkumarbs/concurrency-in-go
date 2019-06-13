package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Work defines the unit/task for the load balancer
type Work struct {
	id, x, y, z int
}

func worker(in <-chan *Work, numWorkers int) (chan *Work, *sync.WaitGroup) {
	var wg sync.WaitGroup
	out := make(chan *Work)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)

		go func(workerId int) {
			defer wg.Done()

			for w := range in {
				fmt.Printf("Worker %v got job: %v \n", workerId, w.id)
				w.z = w.x * w.y
				time.Sleep(time.Duration(w.z) * time.Second)
				fmt.Printf("Worker %v finished job: %v \n", workerId, w.id)
				out <- w
			}
		}(i)
	}

	return out, &wg
}

func enqueueJobs(jobs []*Work) <-chan *Work {
	in := make(chan *Work)

	go func() {
		defer close(in)
		for _, job := range jobs {
			fmt.Printf("Load balancer got job: %v \n", job.id)
			in <- job
		}
	}()
	return in
}

func outputJobs(out <-chan *Work, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for job := range out {
		fmt.Printf("Load balancer finished job: %v \n", job.id)
	}
}

func main() {
	var wg sync.WaitGroup

	workJobs := []*Work{}
	for i := 1; i <= 5; i++ {
		work := Work{
			id: i,
			x:  rand.Intn(5),
			y:  rand.Intn(5),
		}
		workJobs = append(workJobs, &work)
	}

	fmt.Printf("\nInit Jobs \n")
	for _, job := range workJobs {
		fmt.Printf("Job: %+v \n", *job)
	}
	fmt.Printf("\n")

	in := enqueueJobs(workJobs)
	out, workerWg := worker(in, 5)

	go outputJobs(out, &wg)

	workerWg.Wait()
	close(out)

	wg.Wait()

	fmt.Printf("\nProcessed Jobs \n")
	for _, job := range workJobs {
		fmt.Printf("Job: %+v \n", *job)
	}
	fmt.Printf("\n")
}
