package main

import "fmt"

// Variadic function
func add(nums ...int) int {
    total := 0
    for _, val := range nums {
        total += val
    }
    return total
}

func main() {
    nums := []int{2, 4, 10, 0}

    sum := add(1, 5, 9, 10)  // direct integers
    sum1 := add(nums...)     // unpack slice into arguments

    fmt.Println(sum, sum1)   // Output: 25 16
}
