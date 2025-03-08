package main

import (
	"fmt"
	"sync"
	"time"
)

const maxWorkers = 3

func worker(id int, jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	jobQueue := make(chan int, 10)

	// Worker を3つ起動
	for i := 1; i <= maxWorkers; i++ {
		wg.Add(1)
		go worker(i, jobQueue, &wg)
	}

	// ジョブを追加
	for j := 1; j <= 10; j++ {
		jobQueue <- j
	}
	close(jobQueue)

	wg.Wait()
	fmt.Println("All jobs completed.")
}