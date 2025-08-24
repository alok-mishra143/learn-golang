package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	// --------------------------
	// 1. Basic switch
	// --------------------------
	switch i {
	case 1:
		fmt.Println("this is 1")
	case 2:
		fmt.Println("this is 2")
	default:
		fmt.Println("this is default one")
	}

	// --------------------------
	// 2. Multiple conditions in one case
	// --------------------------
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // matches either Saturday OR Sunday
		fmt.Println("this is Weekend ")
	default:
		fmt.Println("this is weekday")
	}

	// --------------------------
	// 3. Type switch
	// --------------------------
	whoami := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("this is int")
		case string:
			fmt.Println("this is string")
		case bool:
			fmt.Println("this is bool")
		default:
			fmt.Println("this is other type", t)
		}
	}

	// test calls
	whoami("hello") // → string
	whoami(true)    // → bool
}
