package main

import (
	"fmt"
	"math/rand"
)

func getRandNumberStream(done chan interface{}) <-chan int {
	randNumberStream := make(chan int)
	go func() {
		defer close(randNumberStream)

		for {
			select {
			case <-done:
				fmt.Println("Closing chan and goroutine")
				return
			case randNumberStream <- rand.Int():
			}
		}
	}()
	return randNumberStream
}

func main() {
	done := make(chan interface{})
	defer close(done)

	randNumberStream := getRandNumberStream(done)

	for i := 0; i < 5; i++ {
		fmt.Println(<-randNumberStream)
	}

}
