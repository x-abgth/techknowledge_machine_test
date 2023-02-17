package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("Print numbers from 1 - 100 concurrently")
	for i := 1; i <= 100; i++ {
		// For each loop we say to wait for the go routine one time
		wg.Add(1)

		// Goroutine to print numbers
		go func(num int) {
			defer wg.Done()
			fmt.Println(num)
		}(i)

		// Wait until the go routine finish the task
		wg.Wait()
	}
}
