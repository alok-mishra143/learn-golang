package main

import "fmt"

func main() {
	/*
		📌 The `range` keyword in Go
		- `range` is used in for-loops to iterate over different data structures.
		- It returns one or two values depending on the type:
		  ✅ For slices/arrays → index and value
		  ✅ For maps          → key and value
		  ✅ For strings       → index and rune (Unicode code point)
		  ✅ For channels      → value (until channel is closed)

		👉 You can ignore values using `_` (underscore).
	*/

	// Example 1: Range over a slice
	nums := []int{1, 2, 3, 4, 5}
	for idx, value := range nums {
		fmt.Println("Index:", idx, "Value:", value)
	}

	// Example 2: Range over a map
	mp := map[string]int{}
	mp["alok"] = 1
	mp["test"] = 2
	mp["ok"] = 3
	mp["lol"] = 44

	for key, value := range mp {
		fmt.Println("Key:", key, "Value:", value)
	}

	// Example 3: Range over a string
	// Gives index and rune (Unicode code point)
	for i, v := range "alok" {
		fmt.Println("Index:", i, "Rune:", v, "Character:", string(v))
	}

	// Example 4: Ignoring values
	for idx := range nums {
		fmt.Println("Only index:", idx)
	}

	for _, value := range nums {
		fmt.Println("Only value:", value)
	}
}
