// WaitGroup is a great way to wait for a set of concurrent operations
// to complete when you either don’t care about the result of the concurrent
// operation, or you have other means of collecting their results.

// Here’s a basic example of using a WaitGroup to wait for goroutines to complete:

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	firstRoutine := func() {
		defer wg.Done()
		fmt.Println("1: Sleeping for 3 seconds")
		time.Sleep(1 * time.Second)
		fmt.Println("1: Done sleeping for 3")
	}

	secondRoutine := func() {
		defer wg.Done()
		fmt.Println("2: Sleeping for 5 seconds")
		time.Sleep(1 * time.Second)
		fmt.Println("2: Done sleeping for 5")
	}

	wg.Add(2)
	go firstRoutine()
	go secondRoutine()
	wg.Wait()
	fmt.Println("All goroutines complete.")
}
