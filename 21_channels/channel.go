package main

import (
	"fmt"
	"sync"
)

// Sum function calculates the sum of two numbers and sends it to the channel
func Sum(total chan int, n1 int, n2 int) {
	sum := n1 + n2
	total <- sum // send the result to the channel
}

func main() {

	// --------------------------------------------------
	// Channels in Go:
	// A channel is a way for goroutines to communicate with each other.
	// You can send data into a channel and receive data from it.
	// Channels can be:
	// 1. Unbuffered: The sender waits until the receiver receives the value.
	// 2. Buffered: The sender can send multiple values without waiting, up to the buffer size.
	// --------------------------------------------------

	// Example of an unbuffered channel
	makeChan := make(chan int) // unbuffered channel

	// Start Sum in a goroutine
	go Sum(makeChan, 10, 15)

	// Receive the result from the channel
	fmt.Println("The sum is", <-makeChan)

	// Close the channel after receiving (good practice)
	close(makeChan)

	// --------------------------------------------------
	// Buffered Channels
	// A buffered channel can hold a limited number of values before the sender is blocked.
	// --------------------------------------------------
	bufferedChan := make(chan int, 3) // buffer size of 3
	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3
	fmt.Println("Buffered channel values:", <-bufferedChan, <-bufferedChan, <-bufferedChan)

	// --------------------------------------------------
	// WaitGroup in Go:
	// WaitGroup is used to wait for a collection of goroutines to finish.
	// Unlike channels, it does not carry data; it only tracks completion.
	// --------------------------------------------------
	var wg sync.WaitGroup

	wg.Add(2) // we have 2 goroutines to wait for

	go func() {
		defer wg.Done() // mark this goroutine as done when finished
		fmt.Println("Goroutine 1 done")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2 done")
	}()

	wg.Wait() // wait for all goroutines in the WaitGroup to finish
	fmt.Println("All goroutines finished")

	// --------------------------------------------------
	// Summary:
	// - Channel: for passing data between goroutines.
	// - Buffered channel: allows sending multiple values without waiting for a receiver immediately.
	// - WaitGroup: for waiting until goroutines finish, but does NOT pass data.
	// --------------------------------------------------
}
