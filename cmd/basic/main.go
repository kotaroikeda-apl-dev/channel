package main

import "fmt"

func main() {
	ch := make(chan string)

	// Goroutine でデータを送信
	go func() {
		ch <- "Hello from Goroutine!"
	}()

	// データを受信
	msg := <-ch
	fmt.Println(msg)
}