package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan int) {
	for v := range ch {
		fmt.Printf("worker %d received %d\n", id, v)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int)
	go worker(1, ch)
	go worker(2, ch)

	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	// 等待 goroutine 完成
	time.Sleep(500 * time.Millisecond)
}
