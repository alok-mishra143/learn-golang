package main

import (
	"fmt"
	"maps" // Go 1.21+ provides this package for map utilities
)

func main() {
	// -----------------------------
	// Creating maps
	// -----------------------------

	// Way 1: Using make()
	mp := make(map[string]int)

	// Way 2: Using map literal
	mp1 := map[string]int{}

	// -----------------------------
	// Adding values to a map
	// -----------------------------
	mp["alok"] = 1
	mp["lol"] = 2

	// -----------------------------
	// Getting values
	// -----------------------------
	fmt.Println("alok:", mp["alok"])       // existing key → returns value
	fmt.Println("mishra:", mp["mishra"])   // non-existing key → returns zero value (0 for int)

	// -----------------------------
	// Deleting a key
	// -----------------------------
	delete(mp, "alok")
	fmt.Println("alok after delete:", mp["alok"]) // returns 0 (since key no longer exists)

	// -----------------------------
	// Clearing a map
	// -----------------------------
	clear(mp)
	fmt.Println("after clear:", mp) // map is now empty

	// -----------------------------
	// Checking if a key exists
	// -----------------------------
	mp1["hello"] = 1
	mp1["test"] = 2
	mp1["not"] = 0

	// The 2-value assignment tells us if key exists
	_, ok := mp1["not"]

	if ok {
		fmt.Println("Key exists in map")
	} else {
		fmt.Println("Key does not exist")
	}

	// -----------------------------
	// Comparing two maps (Go 1.21+)
	// -----------------------------
	fmt.Println("Are maps equal?:", maps.Equal(mp, mp1))
}
