package main

import "fmt"

func main() {
	// for 经典形式
	for i := 0; i < 3; i++ {
		fmt.Println("for:", i)
	}

	// 类似 while
	n := 0
	for n < 3 {
		fmt.Println("while-like:", n)
		n++
	}

	// if / else
	if n == 3 {
		fmt.Println("if: n == 3")
	} else {
		fmt.Println("else")
	}

	// switch
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}
}
