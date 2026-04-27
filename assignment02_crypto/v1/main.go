package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

// 1. 定义统一的接口
type CryptoTask interface {
	Name() string
	Execute(data string) string
}

// 2. 实现具体功能：字符串反转
type ReverseTask struct{}

func (t ReverseTask) Name() string { return "String Reverse" }
func (t ReverseTask) Execute(data string) string {
	runes := []rune(data)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 3. 实现具体功能：MD5摘要
type MD5Task struct{}

func (t MD5Task) Name() string { return "MD5 Hash" }
func (t MD5Task) Execute(data string) string {
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// 4. 实现具体功能：SHA256摘要
type SHA256Task struct{}

func (t SHA256Task) Name() string { return "SHA-256 Hash" }
func (t SHA256Task) Execute(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func main() {
	input := "hello blockchain"

	// 通过切片统一管理接口实例
	tasks := []CryptoTask{
		ReverseTask{},
		MD5Task{},
		SHA256Task{},
	}

	fmt.Println("--- V1: 顺序执行基础任务 ---")
	for _, task := range tasks {
		result := task.Execute(input)
		fmt.Printf("[%s] Input: %s -> Output: %s\n", task.Name(), input, result)
	}
}
