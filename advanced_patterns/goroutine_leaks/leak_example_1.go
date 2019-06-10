package main

import (
	"fmt"
	"time"
)

func leak() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited")
			defer close(completed)

			for s := range strings {
				fmt.Println("Got: ", s)
			}
		}()

		return completed
	}

	doWork(nil)
	fmt.Println("Done!")
}

func fixedLeak() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited")
			defer close(terminated)

			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()

		return terminated
	}

	done := make(chan interface{})
	strings := make(chan string)
	terminated := doWork(done, strings)

	go func() {
		strings <- "sanath"
		time.Sleep(3 * time.Second)
		fmt.Println("Cancelling doWork goroutine..")
		close(done)
	}()

	<-terminated
	fmt.Println("Done!")
}

func main() {
	// leak()
	fixedLeak()
}
