// Multiplexing / Fan In

package main

import (
	"fmt"
	"time"
)

func slow(done chan interface{}) <-chan string {
	slowStream := make(chan string)
	msg := "I am slow"

	go func() {
		defer close(slowStream)

		for {
			select {
			case <-done:
				return
			case slowStream <- msg:
				time.Sleep(3 * time.Second)
			}
		}
	}()

	return slowStream
}

func fast(done chan interface{}) <-chan string {
	fastStream := make(chan string)
	msg := "I am fast"

	go func() {
		defer close(fastStream)

		for {
			select {
			case <-done:
				return
			case fastStream <- msg:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return fastStream
}

func fanIn(done chan interface{}, slowStream, fastStream <-chan string) <-chan string {
	fanInStream := make(chan string)
	go func() {
		for {
			select {
			case <-done:
				return
			case val := <-slowStream:
				fanInStream <- val
			case val := <-fastStream:
				fanInStream <- val
			}
		}
	}()
	return fanInStream
}

func main() {
	done := make(chan interface{})
	defer close(done)

	fanInStream := fanIn(done, slow(done), fast(done))
	for i := 0; i < 10; i++ {
		fmt.Println(<-fanInStream)
	}
}
