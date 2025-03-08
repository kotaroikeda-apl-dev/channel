package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs chan int, results chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		time.Sleep(2 * time.Second)
		results <- fmt.Sprintf("Worker %d finished job %d", id, job)
	}
}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 10)
	results := make(chan string, 10)

	// Worker を3つ起動
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// ジョブを追加
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	// 結果を出力
	for result := range results {
		fmt.Println(result)
	}
}
