package main

import (
	"fmt"
	"time"
)

func patternA() {
	doneChan := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		close(doneChan)
	}()

	workcount := 0

	// A default statement allows you to exit a select block
	// without blocking. Usually you’ll see a default clause
	// used in conjunction with a for-select loop.

	// This allows a goroutine to make progress on work while
	// waiting for another goroutine to report a result.

	// Here’s an example of that:

loop:
	for {
		select {
		case <-doneChan:
			break loop
		default:
		}

		// Simulate work
		workcount++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workcount)
}

func patternB() {
	doneChan := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		close(doneChan)
	}()

	workcount := 0

	// A default statement allows you to exit a select block
	// without blocking. Usually you’ll see a default clause
	// used in conjunction with a for-select loop.

	// This allows a goroutine to make progress on work while
	// waiting for another goroutine to report a result.

	// Here’s an example of that:
	for {
		select {
		case <-doneChan:
			fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workcount)
			return
		default:
			// Simulate work
			workcount++
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	patternB()
}
