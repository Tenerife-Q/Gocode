package main

import "fmt"

// 演示常用关键字的基本用法（只做示例，避免进入并发复杂细节）

const Pi = 3.14 // const

type MyInt int // type

func deferExample() {
	defer fmt.Println("defer: 在函数返回前执行")
	fmt.Println("deferExample: 正在运行")
}

func main() {
	// package, import, func 已在文件顶部体现

	// var / 短声明
	var a int = 10
	b := 20

	// if / else / return
	if a < b {
		fmt.Println("if: a < b")
	} else {
		fmt.Println("else")
	}

	// for (见另一个示例文件) 这里演示 range
	s := []string{"go", "rust", "python"}
	for i, v := range s {
		fmt.Println("range", i, v)
	}

	// fallthrough 在 switch 中使用
	switch a {
	case 10:
		fmt.Println("case 10: a == 10")
		fallthrough
	case 11:
		fmt.Println("case 11: （来自 fallthrough）")
	default:
		// do nothing
	}

	// defer
	deferExample()

	// break / continue 在循环中使用（见 loops 示例）

	// iota（在 const 中常用）
	const (
		Sunday = iota
		Monday
		Tuesday
	)
	fmt.Println("iota example:", Sunday, Monday, Tuesday)

	// type/struct 在其他文件示例中展示
	var x MyInt = 5
	fmt.Println("MyInt:", x)

	// return 在函数中使用（见 deferExample）
}
