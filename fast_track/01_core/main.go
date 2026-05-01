package main

import (
	"errors"
	"fmt"
)

/*
=============================================================
🌟 运行与编译指南：
1. 直接运行 (最常用作本地测试)：go run main.go
2. 编译为可执行文件 (常用于部署)：go build -o myapp.exe main.go
   -> 编译后你会得到一个独立的 myapp.exe，可以在任何没装 Go 的机器上跑。
=============================================================
*/

// ==========================================
// 1. 函数多返回值 & 错误处理 (Go 最重要的基操)
// 举一反三：在 Web 后端中，去数据库查用户，查不到就会返回 (nil, error)
// ==========================================
func findUser(id int) (string, error) {
	// 模拟数据库里只有 id:1 的用户
	if id == 1 {
		return "Alice", nil // nil 表示没有错误，安全通过
	}
	return "", errors.New("用户不存在 (User Not Found)")
}

func main() {
	fmt.Println("--- 1. 函数多返回值与错误处理 ---")
	// 经典范式：先查值，再用 if err != nil 拦截错误
	name, err := findUser(2)
	if err != nil {
		fmt.Printf(" 报错拦截: %v\n", err)
	} else {
		fmt.Printf(" 找到用户: %s\n", name)
	}

	fmt.Println("\n--- 2. 切片(Slice): 动态数组 ---")
	// ==========================================
	// 2. 切片 (Slice)
	// 用处：当你不知道要存多少个元素时用它。比定长数组(Array)灵活一百倍。
	// 举一反三：你可以用它来临时存储页面上拉取的 100 条商品信息。
	// ==========================================
	users := []string{"Bob", "Charlie"} // 初始 2 个元素
	users2 := make([]string, 0, 10)     // 预分配容量为 10 的切片，初始长度为 0
	fmt.Printf("初始用户列表: %v\n", users)
	fmt.Printf("预分配切片: %v (len=%d, cap=%d)\n", users2, len(users2), cap(users2))
	users = append(users, "Dave")         // 追加 1 个，变成 3 个
	users = append(users, "Eve", "Frank") // 一次性追加多个

	// 切片的“切法”操作 [开始索引 : 结束索引] (左闭右开)
	fmt.Printf("所有用户: %v\n", users)
	fmt.Printf("取前两个用户: %v\n", users[0:2]) // 包含索引 0, 1
	fmt.Printf("取从索引 2 到结尾的用户: %v\n", users[2:])

	fmt.Println("\n--- 3. 字典(Map): 键值对 ---")
	// ==========================================
	// 3. 字典 (Map)
	// 用处：实现 O(1) 的超快查找。例如缓存用户的在线状态，或者记录成绩。
	// 举一反三：JSON 解析后的复合数据，很多时候都会转成 Map 结构处理。
	// ==========================================
	userScores := map[string]int{
		"Alice": 99,
		"Bob":   82,
	}
	userScores["Charlie"] = 75 // 增加新键值对
	delete(userScores, "Bob")  // 删除键值对

	// 🔥 重要特性：Map 查值判断键存不存在 (" comma ok " 语法)
	score, exists := userScores["Bob"]
	if exists {
		fmt.Printf("Bob的成绩是: %d\n", score)
	} else {
		fmt.Println("Bob不在成绩单中(被开除了)！")
	}

	fmt.Println("\n--- 4. 循环控制 (for 与 range) ---")
	// ==========================================
	// 4. range 循环语句
	// 用处：最优雅的遍历方式，专门对付 Slice, Map 和字符串。
	// 举一反三：在你的 fitness_task_manager 中，可以用它来遍历当天所有的健身任务。
	// ==========================================
	fmt.Println("遍历数组切片：")
	for i, u := range append(users, "Grace") { // 可以在遍历时临时 append
		// _ 表示丢弃不用。比如写成：for _, u := range users
		fmt.Printf("  第 %d 号人员是 %s\n", i, u)
	}

	fmt.Println("遍历字典 Map：")
	for key, val := range userScores {
		fmt.Printf("  姓名: %s, 成绩: %d\n", key, val)
	}
}
