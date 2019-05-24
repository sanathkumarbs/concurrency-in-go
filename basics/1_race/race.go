// Use go run -race race.go to see how go detects Data Race

package main

import (
	"fmt"
	"time"
)

var data int

func main() {
	// Let's start a goroutine
	go func() {
		data++
	}()

	// Let's check what the data is print accordingly
	// This will be non-deterministic beacause of race

	if data == 0 {
		// One would expect data to be '0' because of the if condition being true
		time.Sleep(2 * time.Second)

		// But, meanwhile it could have gotten updated because of the goroutine
		fmt.Printf("Expected data to be 0, got %v \n", data)
	} else {
		fmt.Printf("Expected data to be not 0, got %v \n", data)
	}
}
