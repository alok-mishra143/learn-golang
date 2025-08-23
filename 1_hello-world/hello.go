package main  // Every Go program starts with a package declaration.
// "main" is the entry point package for executable programs.

import "fmt"  // Importing the "fmt" package, which provides functions for formatted I/O
// (like printing text to the console).

func main() { // The main function is where execution starts in a Go program.
	
    // Println prints the text with a newline at the end
    fmt.Println("Hello, World!")  // Output: Hello, World!

    // Print prints the text without adding a newline
    fmt.Print("hello")            // Output: hello (cursor stays on same line)

    // Printf allows formatting text (f = format).
    // In this case, no format specifiers are used, so it just prints the string.
    fmt.Printf("this is print")   // Output: this is print
}
