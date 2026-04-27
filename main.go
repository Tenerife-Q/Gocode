package main // 声明这是主包，只有 main 包才能生成可执行文件

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(" Go 项目启动成功！")
	fmt.Printf("当前时间是: %s\n", time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println("正在准备起飞...")
	time.Sleep(1 * time.Second)

	fmt.Println("Hello, Gopher! 你的环境已经完美打通。")

	// 基础输入输出练习
	var name string
	fmt.Print("请输入你的名字: ")
	fmt.Scanln(&name)
	fmt.Printf("你好, %s! 欢迎学习 Go 语言。\n", name)

	// 变量和运算练习
	var a, b int
	fmt.Print("请输入两个整数（以空格分隔）: ")
	fmt.Scan(&a, &b)
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	if b != 0 {
		fmt.Printf("%d / %d = %d\n", a, b, a/b)
	} else {
		fmt.Println("除数不能为零！")
	}

	// 简单算法练习
	// 1. 计算数组的最大值
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	maxVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}
	fmt.Printf("数组 %v 的最大值是: %d\n", arr, maxVal)

	// 2. 判断一个数是否是素数
	var num int
	fmt.Print("请输入一个正整数: ")
	fmt.Scan(&num)
	isPrime := true
	if num <= 1 {
		isPrime = false
	} else {
		for i := 2; i <= num/i; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
	}
	if isPrime {
		fmt.Printf("%d 是一个素数。\n", num)
	} else {
		fmt.Printf("%d 不是一个素数。\n", num)
	}

	// 3. 计算斐波那契数列的前 n 项
	var n int
	fmt.Print("请输入斐波那契数列的项数: ")
	fmt.Scan(&n)
	fib := make([]int, n)
	if n > 0 {
		fib[0] = 0
	}
	if n > 1 {
		fib[1] = 1
	}
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	fmt.Printf("斐波那契数列的前 %d 项是: %v\n", n, fib)

	// 冒泡排序算法
	// 修复变量名冲突问题
	bubbleSortArr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("原始数组: %v\n", bubbleSortArr)
	for i := 0; i < len(bubbleSortArr)-1; i++ {
		for j := 0; j < len(bubbleSortArr)-i-1; j++ {
			if bubbleSortArr[j] > bubbleSortArr[j+1] {
				// 交换元素
				bubbleSortArr[j], bubbleSortArr[j+1] = bubbleSortArr[j+1], bubbleSortArr[j]
			}
		}
	}
	fmt.Printf("排序后数组: %v\n", bubbleSortArr)
}
