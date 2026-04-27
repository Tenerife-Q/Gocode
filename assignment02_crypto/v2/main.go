package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
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

// 【新增功能 1】：Base64编码
type Base64Task struct{}

func (t Base64Task) Name() string { return "Base64 Encode" }
func (t Base64Task) Execute(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// 【新增功能 2】：Hex十六进制编码
type HexTask struct{}

func (t HexTask) Name() string { return "Hex Encode" }
func (t HexTask) Execute(data string) string {
	return hex.EncodeToString([]byte(data))
}

func main() {
	input := "hello blockchain"

	// 直接在任务列表中追加，无需修改调用逻辑
	tasks := []CryptoTask{
		ReverseTask{},
		MD5Task{},
		SHA256Task{},
		Base64Task{},
		HexTask{},
	}

	fmt.Println("--- V2: 迭代新增任务 ---")
	for _, task := range tasks {
		result := task.Execute(input)
		fmt.Printf("[%s] Output: %s\n", task.Name(), result)
	}
}
