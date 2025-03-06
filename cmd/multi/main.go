package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- fmt.Sprintf("Worker %d finished", id)
}

func main() {
	ch := make(chan string)

	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	// 結果を受信
	for i := 1; i <= 3; i++ {
		fmt.Println(<-ch)
	}
}
