package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("ワーカー: %d; 処理中...\n", id)
	time.Sleep(time.Second)
	fmt.Printf("ワーカー: %d; 完了\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
