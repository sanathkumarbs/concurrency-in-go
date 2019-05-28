package main

import (
	"fmt"
	"sync"
	"time"
)

const runtime = 1 * time.Second

func greedyWorker(wg *sync.WaitGroup, sharedLock *sync.Mutex) {
	defer wg.Done()

	var count int
	for begin := time.Now(); time.Since(begin) <= runtime; {
		sharedLock.Lock()
		time.Sleep(3 * time.Nanosecond)
		sharedLock.Unlock()
		count++
	}

	fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
}

func politeWorker(wg *sync.WaitGroup, sharedLock *sync.Mutex) {
	defer wg.Done()

	var count int
	for begin := time.Now(); time.Since(begin) <= runtime; {
		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		count++
	}

	fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
}

func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex

	wg.Add(2)
	go greedyWorker(&wg, &sharedLock)
	go politeWorker(&wg, &sharedLock)
	wg.Wait()
}
