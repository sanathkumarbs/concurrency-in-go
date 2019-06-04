package main

import (
	"fmt"
	"sync"
	"time"
)

var queue = make([]interface{}, 0, 10)

func removeFromQueue(delay time.Duration, cond *sync.Cond) {
	time.Sleep(delay)

	cond.L.Lock()
	queue = queue[1:]
	fmt.Println("Removed from queue")
	cond.L.Unlock()
	cond.Signal()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})

	for i := 0; i < 10; i++ {
		cond.L.Lock()

		for len(queue) == 2 {
			cond.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1*time.Second, cond)
		cond.L.Unlock()
	}
}
