package main

import (
	"fmt"
	"os"
)

func main() {
	fname := "examples/12_io/demo.txt"
	data := []byte("Hello from file I/O example\n")

	// 写文件（覆盖）
	if err := os.WriteFile(fname, data, 0644); err != nil {
		fmt.Println("write error:", err)
		return
	}
	fmt.Println("wrote file:", fname)

	// 读文件
	b, err := os.ReadFile(fname)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}
	fmt.Println("read content:\n", string(b))

	// 删除文件（清理）
	_ = os.Remove(fname)
}
