package main

import "fmt"

// 1. 结构体 (代替传统 OOP 中的 Class)
type User struct {
	Name string
	Age  int
}

// 2. 方法 (通过接收者 receiver 绑定到结构体)
// 敲黑板：传指针 (*User) 才能修改真实结构体里的值，并且性能更好，拷贝开销小
func (u *User) GrowUp() {
	u.Age++
}

// 3. 接口 (鸭子类型: 只要实现了SayHello()方法，就是Greeter)
// 考试/项目常考点：解耦神器
type Greeter interface {
	SayHello()
}

// User 隐式实现了 Greeter 接口 (不需要像 Java 那样 implements)
func (u *User) SayHello() {
	fmt.Printf("Hello, I am %s, %d years old.\n", u.Name, u.Age)
}

// 接收接口类型的通用函数
func GreetSomeone(g Greeter) {
	g.SayHello()
}

func main() {
	u := User{Name: "Alice", Age: 20}
	u.GrowUp() // 年龄 +1

	// 指针 &u 被传入，因为它拥有 SayHello() 和 GrowUp() 方法
	GreetSomeone(&u)
}
