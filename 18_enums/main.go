// Program demonstrates usage of custom types and constants in Go.
// It defines an `orderStatus` type with different states and
// provides a function to print the status of an order.

package main

import "fmt"

// orderStatus represents the possible states of an order.
type orderStatus int

// Constants representing various order statuses.
// Using iota automatically increments values starting from 0.
const (
	Received orderStatus = iota // Order has been received
	Confirmed                   // Order has been confirmed
	Cancelled                   // Order has been cancelled
)

// makeOrder prints the current status of the order.
func makeOrder(status orderStatus) {
	fmt.Println("The status is:", status)
}

func main() {
	fmt.Println("Program started...")

	// Example usage: passing `Received` status to makeOrder.
	makeOrder(Received)
}
