package main

import (
	"errors"
	"fmt"
)

// 1. 函数多返回值 & 错误处理 (Go特色)
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0") // 错误作为值返回
	}
	return a / b, nil // nil 代表没有错误
}

func main() {
	// 2. 切片(Slice) - 动态数组，最常用
	names := []string{"Alice", "Bob"}
	names = append(names, "Charlie")

	// 3. 字典(Map) - 键值对
	scores := map[string]int{"Alice": 90, "Bob": 80}
	scores["Charlie"] = 95

	// 4. 循环神器 range (遍历Slice和Map)
	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
	}

	for key, value := range scores {
		fmt.Printf("%s got %d\n", key, value)
	}

	// 5. 错误处理标准范式 (考试/代码必写)
	res, err := divide(10, 2)
	if err != nil {
		fmt.Println("出错了:", err)
	} else {
		fmt.Println("10 / 2 =", res)
	}
}
