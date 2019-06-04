package main

import (
	"fmt"
	"sync"
	"time"
)

var count int

func increment(lock *sync.Mutex, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("ID:%v, Incrementing count from %v \n", id, count)
	lock.Lock()
	count++
	lock.Unlock()
	fmt.Printf("ID:%v, Incremented count to be %v \n", id, count)
}

func decrement(lock *sync.Mutex, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("ID:%v, Decrementing count from %v \n", id, count)
	lock.Lock()
	count--
	lock.Unlock()
	fmt.Printf("ID:%v, Decrementing count to be %v \n", id, count)
}

func goIncrement(lock *sync.Mutex, wg *sync.WaitGroup) {
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go increment(lock, i, wg)
	}
}

func goDecrement(lock *sync.Mutex, wg *sync.WaitGroup) {
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go decrement(lock, i, wg)
	}
}

func main() {
	start := time.Now()

	var lock sync.Mutex
	var wg sync.WaitGroup

	goIncrement(&lock, &wg)
	goDecrement(&lock, &wg)

	wg.Wait()
	fmt.Println("Arithmetic complete.")

	elapsed := time.Since(start)
	fmt.Printf("Took %s to complete \n", elapsed)
}
