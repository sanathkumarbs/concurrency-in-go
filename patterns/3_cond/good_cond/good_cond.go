package main

import (
	"fmt"
	"sync"
	"time"
)

var count int

func conditionMet(lock *sync.Mutex, cond *sync.Cond) bool {
	lock.Lock()

	defer lock.Unlock()
	if count < 100 {
		return true
	}
	cond.Signal()
	return false
}

func shouldIncrement(lock *sync.Mutex, cond *sync.Cond) bool {
	return conditionMet(lock, cond)
}

func updateValues(wg *sync.WaitGroup, lock *sync.Mutex, cond *sync.Cond) {
	defer wg.Done()
	for shouldIncrement(lock, cond) {

		lock.Lock()

		count++
		fmt.Println("Updated count to be ", count)

		time.Sleep(2 * time.Millisecond)

		lock.Unlock()
	}
}

func goodConditionExample(wg *sync.WaitGroup, lock *sync.Mutex, cond *sync.Cond) {
	defer wg.Done()
	var cycles int

	cond.L.Lock()

	for conditionMet(lock, cond) {
		lock.Lock()
		fmt.Println("Waiting... count is ", count)
		lock.Unlock()
		cycles++

		// Here we wait to be notified that the condition has occurred.
		// This is a blocking call and the goroutine will be SUSPENDED.
		cond.Wait()
	}

	cond.L.Unlock()

	fmt.Println("Waited", cycles, "cycles")

	fmt.Println("Condition met! Doing work")
	time.Sleep(3 * time.Second)
	fmt.Println("Condition met! Completed work")
}

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex

	// Here we instantiate a new Cond.
	// The NewCond function takes in a type that satisfies
	// the sync.Locker interface.

	// This is what allows the Cond type to facilitate coordination
	// with other goroutines in a concurrent-safe way.

	cond := sync.NewCond(&sync.Mutex{})

	start := time.Now()

	wg.Add(2)
	go updateValues(&wg, &lock, cond)
	go goodConditionExample(&wg, &lock, cond)
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Took %s to complete \n", elapsed)
}
