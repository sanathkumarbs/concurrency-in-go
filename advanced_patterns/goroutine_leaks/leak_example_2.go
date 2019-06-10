package main

import (
	"fmt"
	"math/rand"
	"time"
)

func leakWrite() {
	newRandStream := func() <-chan int {
		randStream := make(chan int, 1)

		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)

			for {
				fmt.Println("Writing to randStream")
				randStream <- rand.Int()
			}
		}()

		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")

	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d: %d \n", i, <-randStream)
	}
}

func leakWriteFixed() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)

		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)

			for {
				select {
				case <-done:
					fmt.Println("Got a done signal")
					return
				case randStream <- rand.Int():
					fmt.Println("Writing to randStream...")
				}
			}
		}()

		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)

	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d: %d \n", i, <-randStream)
	}

	close(done)
	time.Sleep(2 * time.Second)
	fmt.Println("Done!")
}

func main() {
	// leakWrite()
	leakWriteFixed()
}
