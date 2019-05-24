// Have you ever been in a hallway walking toward another person?
// She moves to one side to let you pass, but you’ve just done the same.
// So you move to the other side, but she’s also done the same.
// Imagine this going on forever, and you understand livelocks.

package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func takeStep(cadence *sync.Cond) {
	cadence.L.Lock()
	cadence.Wait()
	cadence.L.Unlock()
}

func tryDir(dirName string, dir *int32, out *bytes.Buffer, cadence *sync.Cond) bool {
	fmt.Fprintf(out, " %v", dirName)
	atomic.AddInt32(dir, 1)

	takeStep(cadence)

	if atomic.LoadInt32(dir) == 1 {
		fmt.Fprint(out, ". Success!")
		return true
	}

	takeStep(cadence)
	atomic.AddInt32(dir, -1)
	return false
}

func main() {
	var left, right int32

	// Defining a new cadence condition
	cadence := sync.NewCond(&sync.Mutex{})

	// Starting a goroutine which broadcasts the cadence
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	tryLeft := func(out *bytes.Buffer) bool {
		return tryDir("left", &left, out, cadence)
	}

	tryRight := func(out *bytes.Buffer) bool {
		return tryDir("right", &right, out, cadence)
	}

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() { fmt.Println(out.String()) }()
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}

	var peopleInHallway sync.WaitGroup

	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
}
