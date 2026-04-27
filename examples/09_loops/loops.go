package main

import "fmt"

func main() {
	// 经典 for
	for i := 0; i < 3; i++ {
		fmt.Println("for classic:", i)
	}

	// 类似 while 的写法
	i := 0
	for i < 3 {
		fmt.Println("for while-like:", i)
		i++
	}

	// 无限循环 + break
	count := 0
	for {
		if count >= 3 {
			break
		}
		fmt.Println("infinite loop count:", count)
		count++
	}

	// range 遍历切片
	s := []int{10, 20, 30}
	for idx, val := range s {
		fmt.Println("range slice:", idx, val)
	}

	// range 遍历 map（无序）
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println("range map:", k, v)
	}

	// continue
	for j := 0; j < 5; j++ {
		if j%2 == 0 {
			continue // 跳过偶数
		}
		fmt.Println("odd number:", j)
	}

	// 标签与跳出多层循环
outer:
	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			if a == 1 && b == 1 {
				fmt.Println("breaking outer at", a, b)
				break outer
			}
			fmt.Println("nested:", a, b)
		}
	}
}
