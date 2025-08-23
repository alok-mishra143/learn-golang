package main

import "fmt"

func main() {   

    // -------- String variables --------
    
    var name string = "hello"
    // 👉 Declares a variable named 'name' with type string.
    // 👉 Explicitly specifying the type (string).
    
    var name1 = "hello world"
    // 👉 Declares 'name1' and assigns a string value.
    // 👉 Type is automatically inferred (Go sees "hello world" and knows it's a string).

    // -------- Integer variables --------

    var value int = 1
    // 👉 Explicit type (int).
    
    var value1 = 1
    // 👉 Type inferred as int (since 1 is an integer literal).

    // -------- Short-hand syntax (most common in Go) --------
    
    onlyfans := "khalifa"
    // 👉 Short declaration operator `:=` both declares and assigns.
    // 👉 Type automatically inferred as string.
    // 👉 Works only **inside functions** (like `main`).

	// float 
	var floatval float32=8.6
    

    // -------- Printing values --------
    
    fmt.Println(name)   // Output: hello
    fmt.Println(name1)  // Output: hello world
   fmt.Println(onlyfans,value,value1,floatval)
}
