package main

import (
	"fmt"
	"sync"
	"time"
)

var count int

func conditionMet(lock *sync.Mutex) bool {
	lock.Lock()

	defer lock.Unlock()
	if count < 100 {
		return true
	}
	return false
}

func shouldIncrement(lock *sync.Mutex) bool {
	return conditionMet(lock)
}

func updateValues(wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	for shouldIncrement(lock) {

		lock.Lock()

		count++
		fmt.Println("Updated count to be ", count)

		time.Sleep(2 * time.Millisecond)

		lock.Unlock()
	}
}

func veryBadConditionExample(wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	var cycles int

	// greedy check for conditionMet
	// this would consume all cycles of one core
	for conditionMet(lock) {
		lock.Lock()
		fmt.Println("Waiting... count is ", count)
		lock.Unlock()
		cycles++
	}

	fmt.Println("Waited", cycles, "cycles")

	fmt.Println("Condition met! Doing work")
	time.Sleep(3 * time.Second)
	fmt.Println("Condition met! Completed work")
}

func badConditionExample(wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	var cycles int

	// greedy check for conditionMet
	// better than veryBad one but
	//  you’re artificially degrading performance;
	//  too short, and you’re unnecessarily consuming too much CPU time

	for conditionMet(lock) {
		lock.Lock()
		fmt.Println("Waiting... count is ", count)
		lock.Unlock()
		cycles++
		time.Sleep(5 * time.Millisecond)
	}

	fmt.Println("Waited", cycles, "cycles")

	fmt.Println("Condition met! Doing work")
	time.Sleep(3 * time.Second)
	fmt.Println("Condition met! Completed work")
}

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex

	start := time.Now()

	wg.Add(2)
	go updateValues(&wg, &lock)
	go veryBadConditionExample(&wg, &lock)
	// go badConditionExample(&wg, &lock)
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Took %s to complete \n", elapsed)
}
