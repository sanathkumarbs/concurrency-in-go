// A race condition occurs when two or more operations must
// execute in the correct order, but the program has not
// been written so that this order is guaranteed to be maintained.

// Most of the time, this shows up in what’s called a data race,
// where one concurrent operation attempts to read a variable
// while at some undetermined time another concurrent operation
// is attempting to write to the same variable.

// Use go run -race race.go to see how go detects Data Race

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var data int

func main() {
	// Let's start a goroutine
	go func() {
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		data++
	}()

	// Let's check what the data is print accordingly
	// This will be non-deterministic because of race

	if data == 0 {
		// One would expect data to be '0' because of the if condition being true
		time.Sleep(1 * time.Second)

		// But, meanwhile it could have gotten updated because of the goroutine
		fmt.Printf("Expected data to be 0, got %v \n", data)
	} else {
		fmt.Printf("Expected data to be not 0, got %v \n", data)
	}
}
