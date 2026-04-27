package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sync"
)

type CryptoTask interface {
	Name() string
	Execute(data string) string
}

type ReverseTask struct{}

func (t ReverseTask) Name() string { return "String Reverse" }
func (t ReverseTask) Execute(data string) string {
	runes := []rune(data)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type MD5Task struct{}

func (t MD5Task) Name() string { return "MD5 Hash" }
func (t MD5Task) Execute(data string) string {
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}

type SHA256Task struct{}

func (t SHA256Task) Name() string { return "SHA-256 Hash" }
func (t SHA256Task) Execute(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

type Base64Task struct{}

func (t Base64Task) Name() string { return "Base64 Encode" }
func (t Base64Task) Execute(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

type HexTask struct{}

func (t HexTask) Name() string { return "Hex Encode" }
func (t HexTask) Execute(data string) string {
	return hex.EncodeToString([]byte(data))
}

func main() {
	input := "hello blockchain"

	// 并发执行任务集
	concurrentTasks := []CryptoTask{
		MD5Task{},
		SHA256Task{},
	}

	fmt.Println("--- V3: 并发执行优化 ---")

	var wg sync.WaitGroup

	for _, task := range concurrentTasks {
		wg.Add(1)
		// 闭包中传递 task 以避免循环变量捕获问题
		go func(t CryptoTask) {
			defer wg.Done()
			result := t.Execute(input)
			// 注意：并发输出的顺序是不可预测的
			fmt.Printf("[Concurrent - %s] Output: %s\n", t.Name(), result)
		}(task)
	}

	wg.Wait() // 阻塞主协程，直到所有子协程完成
	fmt.Println("所有并发任务执行完毕。")
}
