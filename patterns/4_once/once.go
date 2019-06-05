// As the name implies, sync.Once is a type that utilizes some sync
// primitives internally to ensure that only one call to Do ever
// calls the function passed in—even on different goroutines.

package main

import (
	"fmt"
	"sync"
)

// Count keeps track of count
var Count int

// Increment Count
func Increment() { Count++ }

// Decrement Count
func Decrement() { Count-- }

// Once shows an example of sync.Once
func Once() {
	var wg sync.WaitGroup
	var once sync.Once

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(Increment)
		}()
	}

	wg.Wait()
	fmt.Printf("Count is %d\n", Count)
}

func main() {
	// Once()
	Gotcha()
}
