// A good channel pattern is where you clearly distingush between:
// An Owner, Consumer(s) of channels

// An Owner: Has a write-access view into the channel
// Example: (chan or chan <-)

// Consumer(s): Has ONLY read-access view into the channel
// Example: (<-chan)

package main

import (
	"fmt"
	"sync"
)

// The goroutine that owns a channel should:
// 1. Instantiate the channel.
// 2. Perform writes, or pass ownership to another goroutine.
// 3. Close the channel.
// 4. Encapsulate the previous three things in this list and expose them via a reader channel.

func stringStreamOwner() <-chan string {
	stringStream := make(chan string, 5)

	go func() {
		defer close(stringStream)
		for i := 1; i <= 5; i++ {
			stringStream <- fmt.Sprintf("Hello %v", i)
		}
	}()

	return stringStream
}

// The consumer of a channel should only be concerned about:
// 1. Knowing when a channel is closed.
// 2. Responsibly handling blocking for any reason.

func stringStreamSubscriber(stringStream <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for item := range stringStream {
		fmt.Println(item)
	}
}

func main() {
	var wg sync.WaitGroup

	stringStream := stringStreamOwner()

	wg.Add(1)
	go stringStreamSubscriber(stringStream, &wg)
	wg.Wait()
}
