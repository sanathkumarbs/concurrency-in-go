package main

import "fmt"

func main() {

	generator := func(done chan interface{}, values ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)

			for _, value := range values {
				select {
				case <-done:
					return
				case intStream <- value:
				}
			}
		}()
		return intStream
	}

	add := func(done chan interface{}, intStream <-chan int, additive int) <-chan int {
		addStream := make(chan int)
		go func() {
			defer close(addStream)

			for value := range intStream {
				select {
				case <-done:
					return
				case addStream <- value + additive:
				}
			}
		}()
		return addStream
	}

	multiply := func(done chan interface{}, intStream <-chan int, multiplier int) <-chan int {
		multiplyStream := make(chan int)
		go func() {
			defer close(multiplyStream)

			for value := range intStream {
				select {
				case <-done:
					return
				case multiplyStream <- value * multiplier:
				}
			}

		}()
		return multiplyStream
	}

	done := make(chan interface{})
	defer close(done)

	values := []int{10, 20, 30, 40, 50}
	intStream := generator(done, values...)
	pipeline := add(done, multiply(done, intStream, 2), 10)

	for value := range pipeline {
		fmt.Println(value)
	}
}
