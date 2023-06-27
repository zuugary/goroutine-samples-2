package main

import (
	"fmt"
	"time"
)

func worker(id int, done chan bool) {
	fmt.Printf("ワーカー: %d; 処理中...\n", id)
	time.Sleep(time.Second)
	fmt.Printf("ワーカー: %d; 完了\n", id)

	done <- true
}

func main() {
	done := make(chan bool, 1)

	go worker(1, done)
	go worker(2, done)

	<-done
	<-done
}
