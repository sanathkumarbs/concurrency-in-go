// Fixing the data race using Mutex

package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex
var data int

func main() {
	// Let's use a mutex Lock such that whenever we are
	// executing a crticial section we acquire our lock

	// Also by sharing the same lock with other critical
	// sections, we are serializing execution of multiple
	// critical sections

	// Using this will create lock contention where threads
	// will get blocked until another thread realeases its lock

	// Let's start a goroutine
	go func() {
		mutex.Lock()
		data++
		mutex.Unlock()
	}()

	mutex.Lock()
	if data == 0 {
		time.Sleep(2 * time.Second)
		fmt.Printf("Expected data to be 0, got %v \n", data)
	} else {
		fmt.Printf("Expected data to be not 0, got %v \n", data)
	}
	mutex.Unlock()
}
