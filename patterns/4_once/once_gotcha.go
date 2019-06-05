// This is because sync.Once only counts the number of times Do is called,
// not how many times unique functions passed into Do are called.

// In this way, copies of sync.Once are tightly coupled to the functions
// they are intended to be called with.

package main

import (
	"fmt"
	"sync"
)

// var count int

// func increment() { count++ }
// func decrement() { count-- }

// Gotcha shows a gotcha for sync.Once
func Gotcha() {
	var once sync.Once

	once.Do(Increment)
	once.Do(Decrement)

	fmt.Printf("Count is %d\n", Count)
}
