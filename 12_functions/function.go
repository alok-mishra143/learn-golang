package main

import "fmt"

/*
📌 Functions in Go
- Functions can return other functions (higher-order functions).
- This allows you to create closures and reusable logic.
*/

// Example 1: Function returning one value
func add(a int, b int) int {
	return a + b
}

// Example 2: Function returning multiple values
func greetings() (string, string, string) {
	return "alok", "hello", "mia"
}

// Example 3: Function returning another function
// test returns a function that takes an int and returns an int
func test() func(int) int {
	// return an anonymous function (closure)
	return func(x int) int {
		return x * x
	}
}

func main() {
	// Calling a single-return function
	val := add(10, 20)
	fmt.Println("Addition:", val)

	// Calling a multi-return function
	name, greet, friend := greetings()
	fmt.Println("Multiple values:", name, greet, friend)

	// Ignoring some return values
	_, onlyGreet, _ := greetings()
	fmt.Println("Only greeting:", onlyGreet)

	// Using function that returns another function
	squareFunc := test()      // test() gives back a function
	result := squareFunc(5)   // call returned function
	fmt.Println("Square:", result)
}
