// Example program demonstrating the use of Go generics.
// It defines a generic function `PrintSlice` that can print
// slices of any data type.

package main

import "fmt"

// PrintSlice is a generic function that accepts a slice of any type
// and prints each element on a new line.
//
// T: represents a type parameter (can be int, string, bool, etc.)
// items: slice of type T
func PrintSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	fmt.Println("Printing integer slice:")
	PrintSlice([]int{1, 2, 3, 4})

	fmt.Println("\nPrinting boolean slice:")
	PrintSlice([]bool{true, false, true})

	fmt.Println("\nPrinting string slice:")
	PrintSlice([]string{"Go", "Generics", "Example"})
}
