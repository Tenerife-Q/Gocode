package main

import (
	"fmt"
	"strings"
)

// Task 定义了一个统一的功能接口
type Task interface {
	Name() string
	Execute() string
}

// ----------------- 功能1：字符串转大写 -----------------
type UpperCaseTask struct {
	text string
}

func (t UpperCaseTask) Name() string {
	return "转换为大写"
}

func (t UpperCaseTask) Execute() string {
	return strings.ToUpper(t.text)
}

// ----------------- 功能2：简单计算（加法） -----------------
type AdditionTask struct {
	a, b int
}

func (t AdditionTask) Name() string {
	return "加法计算"
}

func (t AdditionTask) Execute() string {
	return fmt.Sprintf("%d + %d = %d", t.a, t.b, t.a+t.b)
}

// ----------------- 功能3：欢迎消息生成 -----------------
type WelcomeTask struct {
	user string
}

func (t WelcomeTask) Name() string {
	return "生成欢迎语"
}

func (t WelcomeTask) Execute() string {
	return fmt.Sprintf("欢迎回来，%s！", t.user)
}

func main() {
	// 使用接口切片统一管理不同功能的实例
	tasks := []Task{
		UpperCaseTask{text: "hello ai world"},
		AdditionTask{a: 10, b: 25},
		WelcomeTask{user: "Go开发者"},
	}

	fmt.Println("=== V1: 初始版本（基于接口调用） ===")
	// 无论底层实现多有不同，对外只需要调用 Name() 和 Execute()
	for _, task := range tasks {
		fmt.Printf("任务 [%s] 执行结果: %s\n", task.Name(), task.Execute())
	}
}
