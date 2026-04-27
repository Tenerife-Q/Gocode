package main

import (
	"fmt"
	"sync"
	"time"
)

// 考试必考/Go的灵魂：goroutine 和 channel

// worker 模拟一个耗时任务
func worker(id int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done() // 无论函数如何结束，都会通知 WaitGroup 减1

	time.Sleep(time.Millisecond * 500)          // 模拟业务处理
	ch <- fmt.Sprintf("Worker %d finished", id) // 将结果塞入通道 (channel)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 3) // 带缓冲的通道，容量为 3

	// 开启 3 个并发协程
	for i := 1; i <= 3; i++ {
		wg.Add(1)             // 增加一个等待任务
		go worker(i, &wg, ch) // go 关键字直接开启独立并发任务！(这就是Go的魅力)
	}

	wg.Wait() // 阻塞在此，直到 3 个任务都调用 wg.Done()
	close(ch) // 任务发完，必须关闭通道，否则后面 range 会死锁

	// 读取通道收集结果
	for msg := range ch {
		fmt.Println(msg)
	}
}
