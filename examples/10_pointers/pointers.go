package main

import "fmt"

func main() {
	var x int = 42
	var p *int = &x // 指针保存变量地址
	fmt.Println("x:", x)
	fmt.Println("p (addr):", p)
	fmt.Println("*p (value):", *p)

	*p = 100 // 通过指针修改原值
	fmt.Println("x after *p=100:", x)

	// nil 指针
	var q *int
	if q == nil {
		fmt.Println("q is nil")
	}
}
