// The range keyword—used in conjunction with the for
// statement—supports channels as arguments, and will
// automatically break the loop when a channel is closed.

// This allows for concise iteration over the values on a channel.

// Let’s take a look at an example:

package main

import (
	"fmt"
	"time"
)

func main() {
	intStream := make(chan int)

	// Writing five values into the intStream
	go func() {
		for i := 1; i <= 5; i++ {
			intStream <- i
			time.Sleep(1 * time.Second)
		}
	}()

	// Writing five more values into the intStream
	go func() {
		defer close(intStream)

		time.Sleep(6 * time.Second)

		for i := 6; i <= 10; i++ {
			intStream <- i
			time.Sleep(1 * time.Second)
		}
	}()

	// Reading the values in downstream without
	// having to know how many values are being
	// put on the stream

	go func() {
		for integer := range intStream {
			fmt.Printf("go: %v ", integer)
		}
	}()

	// Blocks until the intStream is closed
	for integer := range intStream {
		fmt.Printf("for: %v ", integer)
	}

	fmt.Printf("\n")

}
