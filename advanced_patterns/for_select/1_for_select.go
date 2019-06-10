package main

import "fmt"

func main() {
	stringStream := make(chan string)
	done := make(chan interface{})

	go func() {
		defer close(stringStream)
		for _, s := range []string{"a", "b", "c"} {
			select {
			case <-done:
				return
			case stringStream <- s:
			}
		}
	}()

	for data := range stringStream {
		fmt.Println(data)
	}
}
