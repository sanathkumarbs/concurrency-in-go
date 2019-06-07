// The select statement is the glue that binds channels
// together; it’s how we’re able to compose channels
// together in a program to form larger abstractions.

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	chanStream1 := make(chan interface{})
	chanStream2 := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		close(chanStream1)
	}()

	go func() {
		time.Sleep(3 * time.Second)
		close(chanStream2)
	}()

	fmt.Println("Blocking on read..")
	select {
	case <-chanStream1:
		fmt.Printf("Unblocked chanStream1 %v later \n", time.Since(start))
	case <-chanStream2:
		fmt.Printf("Unblocked chanStream2 %v later \n", time.Since(start))
	case <-time.After(6 * time.Second):
		fmt.Println("Timed out.")
	}
}
