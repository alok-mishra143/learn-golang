package main

import "fmt"

// change takes an int (value type).
// Go always passes function arguments by value (a copy is made).
// So, changes inside this function won't affect the original variable.
func change(val int) {
	val = 5
	fmt.Println("Inside change(), val is set to:", val)
}

// changeNum takes a pointer to an int (*int).
// A pointer stores the memory address of a variable.
// Using *value, we can directly change the original variable's value.
func changeNum(value *int) {
	*value = 10 // dereference pointer and assign new value
	fmt.Println("Inside changeNum(), value is set to:", *value)
}

func main() {
	nums := 1

	// Pass by value (copy).
	// The function gets a copy of nums, so original nums is not changed.
	change(nums)
	fmt.Println("After calling change(), nums is still:", nums)

	// Printing memory address of nums.
	fmt.Println("The memory address of nums is:", &nums)

	// Pass by reference (pointer).
	// We pass the memory address of nums, so the function can modify it directly.
	changeNum(&nums)
	fmt.Println("After calling changeNum(), nums is now:", nums)
}
