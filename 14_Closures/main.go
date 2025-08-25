package main

import "fmt"

// counter is a function that returns another function (a closure).
// A closure "remembers" the variables from the scope in which it was created.
// In this case, the returned function remembers the variable 'count' even after
// counter() has finished executing.
func counter() func() int {
	count := 0 // local variable inside counter()

	// This is the closure:
	// It captures and modifies the 'count' variable from the outer function.
	return func() int {
		count += 1
		return count
	}
}

func main() {
	// Example of using a closure:

	// inc now holds the closure function returned by counter().
	// Even though counter() has finished, 'count' is still alive
	// because it's captured inside the closure.
	inc := counter()

	// Each call to inc() will increment the same 'count' variable.
	fmt.Println(inc()) // Output: 1
	fmt.Println(inc()) // Output: 2
	fmt.Println(inc()) // Output: 3

	// If we create another closure, it will have its own separate 'count'.
	anotherInc := counter()
	fmt.Println(anotherInc()) // Output: 1 (new counter starts fresh)
	fmt.Println(anotherInc()) // Output: 2

	// Closures are powerful when you want to preserve state
	// without using global variables.
}
