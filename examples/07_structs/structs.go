package main

import "fmt"

// 结构体：数据的组合（不包含方法，避免进入面向对象部分）
type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"张三", 30}
	fmt.Println("struct:", p)
	fmt.Println("字段访问:", p.Name, p.Age)
}
