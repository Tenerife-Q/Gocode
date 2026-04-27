package main

import "fmt"

func main() {
	// 数组
	arr := [3]int{1, 2, 3}
	fmt.Println("array:", arr)

	// 切片
	sl := []int{1, 2, 3, 4}
	fmt.Println("slice:", sl, "len:", len(sl), "cap:", cap(sl))

	sl = append(sl, 5)
	fmt.Println("append:", sl)

	sub := sl[1:3]
	fmt.Println("subslice:", sub)
}
