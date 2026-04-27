package main

import (
	"errors"
	"fmt"
)

// 自定义错误类型（实现 error 接口）
type MyError struct {
	Msg string
}

func (e MyError) Error() string { return "MyError: " + e.Msg }

func mayFail(flag bool) error {
	if !flag {
		return errors.New("simple error: flag is false")
	}
	return nil
}

func main() {
	if err := mayFail(false); err != nil {
		fmt.Println("handled:", err)
	}

	if err := mayFail(true); err != nil {
		fmt.Println("not reached")
	}

	// fmt.Errorf 用法
	err := fmt.Errorf("wrapped: %w", errors.New("inner"))
	fmt.Println(err)

	// 自定义错误
	ce := MyError{"something went wrong"}
	fmt.Println(ce)
}
