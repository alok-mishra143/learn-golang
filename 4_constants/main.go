// constants.go
package main

import "fmt"

func main() {
    // -----------------------------
    // Constant Basics
    // -----------------------------
    const pi = 3.14159
    fmt.Println("Value of pi:", pi)

    // -----------------------------
    // Constants cannot be reassigned
    // -----------------------------
    // pi = 3.14 // ❌ This will cause a compile-time error

    // -----------------------------
    // Multiple constants in a block
    // -----------------------------
    const (
        port = 8080
        host = "localhost"
    )
    fmt.Println("Server running on:", host, ":", port)

    // -----------------------------
    // Untyped constants (more flexible)
    // -----------------------------
    const multiplier = 5
    var result = 10 * multiplier // works with int
    var resultF = 2.5 * multiplier // works with float
    fmt.Println("Result (int):", result)
    fmt.Println("Result (float):", resultF)
}
