package main

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}

	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			// fmt.Printf("Sending %v \n", i)
			c <- token
		}
	}

	reciever := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			// fmt.Printf("Receiving %v \n", i)
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go reciever()
	b.StartTimer()
	close(begin)
	wg.Wait()
}

func main() {
	// benchmarkContextSwitch
}
