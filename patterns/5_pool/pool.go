// Pool is a concurrent-safe implementation of the object pool pattern.
// A the pool pattern is a way to create and make available a fixed number,
// or pool, of things for use.

// It’s commonly used to constrain the creation of things that are expensive
// (e.g., database connections) so that only a fixed number of them are ever
// created, but an indeterminate number of operations can still request access
// to these things.

// In the case of Go’s sync.Pool, this data type can be safely used by
// multiple goroutines

package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	// No object in Pool, creates a new one
	instance1 := myPool.Get()

	// No object in Pool - which is available, creates a new one
	instance2 := myPool.Get()

	myPool.Put(instance1)

	// Reuses the pool member instance1
	myPool.Get()

	myPool.Put(instance2)

}
