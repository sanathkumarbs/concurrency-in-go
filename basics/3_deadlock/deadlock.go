package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mutex sync.Mutex
	value int
}

func printSum(v1, v2 *value, wg *sync.WaitGroup) {
	// mark done once when function finishes
	defer wg.Done()

	// acquire lock on v1
	v1.mutex.Lock()
	defer v1.mutex.Unlock()

	// Using sleep to simulate deadlocks
	time.Sleep(2 * time.Second)

	v2.mutex.Lock()
	defer v2.mutex.Unlock()

	fmt.Printf("%v + %v = %v \n", v1.value, v2.value, v1.value+v2.value)
}

func main() {

	var wg sync.WaitGroup
	var a, b value

	// Let's add two goroutines to be 'waited' to complete
	wg.Add(2)

	// We here first acquire a lock on a and then try to get a lock on b
	go printSum(&a, &b, &wg)

	// Meanwhile, in this goroutine we have already acquired lock on b
	// and then we try to get a
	go printSum(&b, &a, &wg)

	// Both these goroutines are going to sit and wait for the other to release
	// the lock, there by hitting a deadlock condition
	wg.Wait()
}
