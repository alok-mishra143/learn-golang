package main

import "fmt"

func main() {
    age := 18

    // -------------------------------
    // Simple if-else
    // -------------------------------
    if age >= 18 {
        fmt.Println("this is if block")
    } else {
        fmt.Println("this is else block")
    }

    // -------------------------------
    // if - else if - else ladder
    // -------------------------------
    if age >= 10 {
        fmt.Println("this is greater than 10")
    } else if age >= 16 {
        fmt.Println("this is 16+")
    } else {
        fmt.Println("now it is 18+")
    }

	
    // -------------------------------
    // if - else if - else ladder with variable declaeration
    // -------------------------------

	if val:=50; val>60 {
		fmt.Println("graete then 60")
	} else if val>100 {
		fmt.Println("greate then 100")
	} else {
		fmt.Println("this is else block")
	}

    // -------------------------------
    // Loop with condition inside
    // -------------------------------
    for i := range 20 { // range over numbers [0..19]
        if i == 18 {
            fmt.Println("18+ now")
        }
    }
}
