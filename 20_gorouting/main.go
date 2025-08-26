// Example program demonstrating the use of goroutines and sync.WaitGroup in Go.
// Multiple tasks run concurrently, and main waits for all of them to finish.

package main

import (
	"fmt"
	"sync"
)

// task simulates a job by printing its ID, then marks itself done.
func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Running task:", id)
}

func main() {
	var wg sync.WaitGroup

	// Launch 10 tasks concurrently.
	const n = 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go task(i, &wg)
	}

	// Wait for all goroutines to complete before exiting.
	wg.Wait()
	fmt.Println("All tasks completed.")
}
