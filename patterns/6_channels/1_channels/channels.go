// Channels are one of the synchronization primitives in Go.
// While they can be used to synchronize access of the memory,
// they are best used to communicate information between goroutines.

// When using channels, you’ll pass a value into a chan variable,
// and then somewhere else in your program read it off the channel.

package main

import "fmt"

func main() {

	// To declare a unidirectional channel, you’ll simply include the <- operator.

	// RECEIVE CHAN:
	// Place the <- operator on the lefthand side:
	// var receiveDataStream <-chan interface{}
	// receiveDataStream = make(<-chan interface{})

	// SEND CHAN:
	// Place the <- operator on the right side:
	// var sendDataStream chan<- interface{}
	// sendDataStream = make(chan<- interface{})

	// TYPED BIDIRECTIONAL CHAN (Single line initialization):
	stringDataStream := make(chan string)

	// Does this goroutine run before program finishes?
	// Yes, because channels are blocking in nature
	go func() {
		stringDataStream <- "Hello Channels!"
	}()

	val, ok := <-stringDataStream
	fmt.Printf("(%v): %v \n", ok, val)

	// What happens if you try reading from the channel now?
	// val, ok = <-stringDataStream
	// fmt.Printf("(%v): %v \n", ok, val)

	// Blocking is defined as:

	// If you are trying to read data from a channel but channel
	// does not have a value available with it,
	// it blocks the current goroutine and unblocks other in a
	// hope that some goroutine will push a value to the channel.
	// Hence, this read operation will be blocking.

	// Similarly, if you are to send data to a channel,
	// it will block current goroutine and unblock others until
	// some goroutine reads the data from it.
	// Hence, this send operation will be blocking.

	// CLOSED CHANNELS
	close(stringDataStream)

	// What if you try reading from the channel now?
	val, ok = <-stringDataStream
	fmt.Printf("(%v): %v \n", ok, val)
}
