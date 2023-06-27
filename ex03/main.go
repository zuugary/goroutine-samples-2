package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		count int
		wg    sync.WaitGroup
	)

	for i := 0; i < 5000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			count++
		}()
	}

	wg.Wait()

	fmt.Printf("count: %d\n", count)
}
