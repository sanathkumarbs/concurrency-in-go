// Closing a channel is also one of the ways you can
// signal multiple goroutines simultaneously.

// If you have n goroutines waiting on a single channel,
// instead of writing n times to the channel to unblock
// each goroutine, you can simply close the channel.

// Remember in “The sync Package” we discussed using
// the sync.Cond type to perform the same behavior.
// You can certainly use that, but as we’ve discussed,
// channels are composable, so this is my favorite way
// to unblock multiple goroutines at the same time.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	chanStream := make(chan int)

	fmt.Println("Starting goroutines...")

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Started goroutine %v \n", i)
			<-chanStream
			fmt.Printf("Unblocked %v \n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(chanStream)
	wg.Wait()
}
