// loops.go
package main

import "fmt"

func main() {
    // ------------------------------------
    // 1. While-style loop (Go uses 'for')
    // Print numbers from 1 to 3
    // ------------------------------------
    i := 1
    for i <= 3 {
        fmt.Println("Count:", i)
        i = i + 1
    }

    // ------------------------------------
    // 2. Infinite loop (use carefully!)
    // Break after 3 iterations
    // ------------------------------------
    j := 0
    for {
        fmt.Println("Ping")
        j++
        if j == 3 {
            break // exit loop
        }
    }

    // ------------------------------------
    // 3. Classic for loop (like C/Java)
    // Print even numbers less than 6
    // ------------------------------------
    for k := 0; k < 6; k += 2 {
        fmt.Println("Even number:", k)
    }

	
    // ------------------------------------
    // 4. Classic for loop (like C/Java)
    // Print even numbers less than 6
    // ------------------------------------
    for i:=range 6 {

		if i==3 {
			continue
		}
		fmt.Println(i)
	}
    // ------------------------------------
    // 5. Loop over a slice (collection)
    // ------------------------------------
    fruits := []string{"apple", "banana", "mango"}
    for index, fruit := range fruits {
        fmt.Println("Fruit", index, "is", fruit)
    }

    // ------------------------------------
    // 6. Loop over a map (key-value)
    // ------------------------------------
    capitals := map[string]string{
        "India":   "New Delhi",
        "Japan":   "Tokyo",
        "France":  "Paris",
    }
    for country, capital := range capitals {
        fmt.Println("Capital of", country, "is", capital)
    }
}
