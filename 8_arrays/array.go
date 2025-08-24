package main

import "fmt"

func main() {
    // ---------------------------
    // 1. Declare an array of size 6 (all default values → 0)
    // ---------------------------
    var num [6]int
    // num[1] = 10 // would assign 10 at index 1
    fmt.Println(num)       // Output: [0 0 0 0 0 0]
    fmt.Println(len(num))  // Output: 6 (array length)

    // ---------------------------
    // 2. Declare & initialize array
    // ---------------------------
    val := [3]int{1, 2, 3}
    fmt.Println(val)       // Output: [1 2 3]

    // ---------------------------
    // 3. 2D array (matrix)
    // ---------------------------
    var matrix [2][3]int // 2 rows, 3 columns
    matrix[0][0] = 1
    matrix[0][1] = 2
    matrix[0][2] = 3
    matrix[1][0] = 4
    matrix[1][1] = 5
    matrix[1][2] = 6

    fmt.Println("2D Array:", matrix)
    // Output: [[1 2 3] [4 5 6]]
}
