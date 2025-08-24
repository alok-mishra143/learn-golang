package main

import (
	"fmt"
	"slices" // Go 1.21+ has a slices package with helper functions
)

func main() {

	/*
		📌 Slices in Go
		- A slice is a flexible, dynamically-sized view into an array.
		- Unlike arrays, slices don’t have a fixed size.
		- They store:
		  1. Pointer to the underlying array
		  2. Length (current elements in the slice)
		  3. Capacity (total space before needing a new array)

		✅ Key Points:
		- Declare: var s []int  // nil slice
		- Initialize with values: s := []int{1,2,3}
		- Use make: s := make([]int, length, capacity)
		- Can grow using append
		- Copy one slice into another using copy()
	*/

	
	// 1. Declaring an empty slice
	var value []int
	fmt.Println("Empty slice:", value) // []

	// 2. Slices with size using make
	val := make([]int, 2) // length=2, capacity=2
	fmt.Println("Slice with make:", val)

	// 3. Another way to make a slice (literal)
	val1 := []int{}
	fmt.Println("Slice literal:", val1)

	// 4. Copying slices
	src := []int{1, 2, 3}
	dest := make([]int, len(src))
	copy(dest, src) // copy(dest, src) ✅ correct syntax
	fmt.Println("Source slice:", src)
	fmt.Println("Copied slice:", dest)

	// 5. Slice operator
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("nums[0:1]:", nums[0:1]) // [10]
	fmt.Println("nums[0:]:", nums[0:])   // [10 20 30 40 50]
	fmt.Println("nums[:3]:", nums[:3])   // [10 20 30]

	// 6. Using slices package (Go 1.21+)
	num1 := []int{1, 2, 3}
	num2 := []int{1, 2, 3}
	fmt.Println("Slices equal:", slices.Equal(num1, num2)) // true

	// 7. 2D slice (slice of slices)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D slice:", matrix)
}
