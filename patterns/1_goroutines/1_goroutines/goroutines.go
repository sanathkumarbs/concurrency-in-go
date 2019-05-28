package main

import (
	"fmt"
	"sync"
)

// You may notice that the goroutines may never finish
// running because main() might exit before the concurrent
// goroutines "finish" their work
func exitsBeforeGoRoutinesRun() {
	sayHello := func() {
		fmt.Println("Hello")
	}

	go func() {
		fmt.Println("world")
	}()
	// () -> invokes the goroutine

	// for named variables, we can invoke the
	// goroutine by doing go func()
	go sayHello()
}

// You may notice that the goroutines will finish running
// before main() exits. Because of the "WaitGroup" which
// is waiting for the '2' goroutines to mark themself
// wg.Done()
func waitUntilGoRoutinesRun() {
	var wg sync.WaitGroup

	wg.Add(2)
	sayHello := func() {
		defer wg.Done()
		fmt.Println("Hello")
	}

	go func() {
		defer wg.Done()
		fmt.Println("world")
	}()

	// for named variables, we can invoke the
	// goroutine by doing go func()
	go sayHello()
	wg.Wait()
}

// Anonymous Functions
func anonymousFunctions() {
	// Anonymous Function : 1
	sayHello := func() {
		fmt.Println("Hello")
	}

	// Anonymous Function : 2
	go func() {
		fmt.Println("world")
	}()

	// for named variables, we can invoke the
	// goroutine by doing go func()
	go sayHello()
}

// Closure Functions
func closureFunctions() {
	var wg sync.WaitGroup

	state := "hello"

	fmt.Printf("state is now %v \n", state)

	wg.Add(1)
	go func() {
		// This is a closure function
		// It is a type of anonymous function. But, it also shares the "state"
		// A closure is a function that captures the state of the surrounding environment.

		defer wg.Done()

		state = "welcome"
	}()
	wg.Wait()

	// Because of closure and the shared "state", we will notice that the
	// "state" var has gotten updated to "welcome"

	fmt.Printf("state is now %v \n", state)
}

// Let's look at the for loop where the "loop variable" is a reference and not a copy
func overwrittenLoopVariables() {
	var wg sync.WaitGroup

	fmt.Println("Example for overwritten variable")

	// This will overwrite word to always reference to "welcome"
	for _, word := range []string{"hello", "world", "welcome"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(word)
		}()
	}
	wg.Wait()

	fmt.Println("Example with the fix for overwritten variable")
	// The fix, pass the value to the closure goroutine by "value"
	// not a reference to the var "word"
	for _, word := range []string{"hello", "world", "welcome"} {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			fmt.Println(word)
		}(word)
	}
	wg.Wait()
}

func main() {
	// Comment everything but one to see the behavior
	// exitsBeforeGoRoutinesRun()
	// waitUntilGoRoutinesRun()
	// anonymousFunctions()
	// closureFunctions()
	// overwrittenLoopVariables()
}
