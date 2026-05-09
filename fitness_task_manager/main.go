package main

// 引入Go语言的标准库
import (
	"bufio"         // 用于处理带缓冲的I/O读取，比如读取用户的整行输入
	"crypto/sha256" // 密码学库，用于计算SHA-256哈希值（用于防篡改链）
	"encoding/hex"  // 用于将哈希计算后的二进制结果转为十六进制字符串
	"encoding/json" // 用于将Go语言的结构体数据与JSON格式相互转换（持久化核心）
	"fmt"           // 格式化I/O，用于在终端打印输出
	"os"            // 操作系统交互，用于文件读写操作
	"strconv"       // 字符串转换库，用于字符串与数字类型之间的转换
	"strings"       // 字符串处理库，用于去除输入两端的空格等
	"time"          // 时间处理库，用于记录任务创建和截止时间
)

// ================== 数据模型 ==================

// FitnessTask 任务数据结构体
// 结构体(Struct)是Go语言中用于组合多个变量的自定义类型，类似于其他语言的类(Class)
// 后面的 `json:"..."` 是结构体标签(Tag)，告诉 json 库在保存文件时该字段叫什么名字
type FitnessTask struct {
	ID          string    `json:"id"`          // 唯一标识符
	Name        string    `json:"name"`        // 任务名称
	Description string    `json:"description"` // 任务详细描述
	Category    string    `json:"category"`    // 分类扩展字段
	Status      string    `json:"status"`      // 当前状态：待处理、进行中、已完成
	Priority    int       `json:"priority"`    // 优先级整数：1-5
	CreatedAt   time.Time `json:"created_at"`  // 创建时的系统时间
	Deadline    time.Time `json:"deadline"`    // 扩展字段：截止时间
}

// LogBlock 链式日志块结构体 (用于操作追踪与防篡改)
type LogBlock struct {
	Index     int       `json:"index"`     // 区块所在索引(第几个操作)
	Timestamp time.Time `json:"timestamp"` // 发生此操作的时间
	Summary   string    `json:"summary"`   // 做了什么的文本摘要，例如"新增任务A"
	PrevHash  string    `json:"prev_hash"` // 上一次操作生成的哈希值(将所有操作串联起来的关键)
	Hash      string    `json:"hash"`      // 本次操作生成的哈希值
}

// 全局变量声明
var (
	// 切片(Slice)：类似于动态数组，可以随时追加元素。由于还没有学习数据库，先用切片把数据全放在内存里。
	tasks      []FitnessTask
	blockchain []LogBlock

	// 数据持久化保存的文件路径
	taskFile  = "tasks_data.json"
	chainFile = "chain_logs.json"
)

// ================== 数据持久化与链式日志核心机制 ==================

// loadData 程序启动时从文件读取JSON数据到内存
func loadData() {
	// os.ReadFile 会把整个文件读成字节数组 ([]byte)
	if file, err := os.ReadFile(taskFile); err == nil {
		// 返回没有错误(err == nil)说明文件存在，通过 Unmarshal 把字节转回 tasks 切片
		json.Unmarshal(file, &tasks)
	}

	// 读取日志链文件
	if cFile, err := os.ReadFile(chainFile); err == nil {
		json.Unmarshal(cFile, &blockchain)
	} else {
		// 如果读取失败(通常是因为第一次运行还没有文件)，则记录一条创世日志启动系统
		recordLog("系统初始化 - 创建创世操作记录")
	}
}

// saveData 将内存里的数据结构重新持久化保存为硬盘上的文件
func saveData() {
	// MarshalIndent 将结构体变成带缩进、易于人类阅读的JSON格式
	data, _ := json.MarshalIndent(tasks, "", "  ")
	// os.WriteFile 写入文件，0644 表示文件权限(可读写)
	os.WriteFile(taskFile, data, 0644)

	cData, _ := json.MarshalIndent(blockchain, "", "  ")
	os.WriteFile(chainFile, cData, 0644)
}

// calculateHash 根据当前参数计算出这一步操作的唯一指纹(哈希值)
func calculateHash(index int, timestamp time.Time, summary, prevHash string) string {
	// 将各项信息拼接成一个长字符串
	record := strconv.Itoa(index) + timestamp.Format(time.RFC3339) + summary + prevHash
	// 初始化SHA256哈希算法对象
	h := sha256.New()
	// 把长字符串放进去运算
	h.Write([]byte(record))
	// h.Sum 最终得到的是乱码二进制，再通过 hex 转为可以看到的由字母数字组成的十六进制字符串
	return hex.EncodeToString(h.Sum(nil))
}

// recordLog 拦截系统的关键操作并生成哈希区块（相当于操作记录本）
func recordLog(summary string) {
	index := len(blockchain)
	// 第一个区块没有上一个哈希指针，统一设为 64 个 0
	prevHash := "0000000000000000000000000000000000000000000000000000000000000000"
	if index > 0 {
		// 取出 blockchain 数组最后一个元素的 Hash 作为当前元素的 PrevHash
		prevHash = blockchain[index-1].Hash
	}

	timestamp := time.Now()
	// 计算本次操作的指纹
	hash := calculateHash(index, timestamp, summary, prevHash)

	// 把新的操作日志追加进去
	blockchain = append(blockchain, LogBlock{
		Index:     index,
		Timestamp: timestamp,
		Summary:   summary,
		PrevHash:  prevHash,
		Hash:      hash,
	})

	// 每次有了新操作后，立即保存数据文件
	saveData()
}

// verifyChain 遍历检查日志记录有没有被人在外部偷偷修改文本
func verifyChain() bool {
	// 从第1个(由0开始)遍历到最后一个
	for i := 1; i < len(blockchain); i++ {
		prev := blockchain[i-1]
		curr := blockchain[i]

		// 1. 检查链条是否断裂：当前块记录的"上一步哈希" 是否等于 上一块真实的哈希
		if curr.PrevHash != prev.Hash {
			return false
		}
		// 2. 检查内容是否被篡改：用现有内容重新计算哈希，看看和原来存下来的哈希能否对上
		if curr.Hash != calculateHash(curr.Index, curr.Timestamp, curr.Summary, curr.PrevHash) {
			return false
		}
	}
	return true
}

// ================== 业务逻辑处理 ==================

// readInput 封装一个通用的命令行输入读取函数，带回显提示
func readInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin) // 绑定标准输入终端
	// 读一直到用户按下回车键 ('\n')
	input, _ := reader.ReadString('\n')
	// 清除因为回车产生的多余空格和换行符
	return strings.TrimSpace(input)
}

// generateID 生成任务的唯一ID，这里简单使用当前的时间戳
func generateID() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// addTask 新增任务的核心流程
func addTask() {
	name := readInput("请输入任务名称: ")
	desc := readInput("请输入任务描述: ")
	cate := readInput("请输入分类(例如: 学习/生活/锻炼): ")
	priStr := readInput("请输入优先级(1-5整数): ")

	// strconv.Atoi 尝试把字符串转换为整数 (ASCII to Int)
	priority, _ := strconv.Atoi(priStr)

	// 实例化结构体对象
	task := FitnessTask{
		ID:          generateID(),
		Name:        name,
		Description: desc,
		Category:    cate,
		Status:      "待处理",
		Priority:    priority,
		CreatedAt:   time.Now(),
		Deadline:    time.Now().Add(24 * time.Hour), // 默认演示：设定为24小时后截止
	}

	// append 是Go特有的内建函数，将新元素追加到切片末尾
	tasks = append(tasks, task)

	// Sprintf: 不直接打印到屏幕，而是把变量拼成字符串，交给日志系统
	recordLog(fmt.Sprintf("新增任务: [%s] %s", task.ID, task.Name))

	fmt.Println("任务添加成功，并已完成数据持久化！")
}

// queryTasks 查看所有任务
func queryTasks() {
	if len(tasks) == 0 {
		fmt.Println("当前没有任何任务。")
		return
	}
	fmt.Println("\n--- 任务列表 ---")
	// 格式化输出表头： %-13s 表示对齐并在右侧填充至13个字符
	fmt.Printf("%-13s | %-10s | %-6s | %-6s | %-4s | %-10s\n", "ID", "名称", "分类", "状态", "优先级", "创建日期")
	fmt.Println(strings.Repeat("-", 65))

	// for range 是Go语言遍历切片(数组)最常见的语法
	// _ 是索引用不着，所以用下划线忽略；t 表示当前取出来的单条任务数据
	for _, t := range tasks {
		// t.CreatedAt.Format 是Go语言特殊的时间格式化语法 (用固定的 2006-01-02 15:04:05 表示格式模板)
		fmt.Printf("%-13s | %-10s | %-6s | %-6s | %-4d | %-10s\n", t.ID, t.Name, t.Category, t.Status, t.Priority, t.CreatedAt.Format("01-02 15:04"))
	}
}

// updateTask 更新任务的状态字段
func updateTask() {
	id := readInput("请输入要更新状态的任务ID: ")

	// i 是当前元素的索引(0, 1, 2...)
	for i, t := range tasks {
		if t.ID == id {
			newStatus := readInput(fmt.Sprintf("输入新状态 (当前: %s, 建议填: 待处理/进行中/已完成): ", t.Status))
			if newStatus != "" {
				// 修改指定索引所在元素的值
				tasks[i].Status = newStatus
				recordLog(fmt.Sprintf("更新任务 [%s] 的状态为: %s", t.ID, newStatus))
				fmt.Println("任务状态更新成功！")
			}
			return
		}
	}
	fmt.Println("错误：未找到指定的任务ID。")
}

// deleteTask 从切片中移除任务数据
func deleteTask() {
	id := readInput("请输入要删除的任务ID: ")
	for i, t := range tasks {
		if t.ID == id {
			// 在Go中删除切片的某个元素，常用的技巧：取你要删除项的前半段 + 后半段，然后拼接起来
			// tasks[:i] 获取 [0 到 i-1] 的元素
			// tasks[i+1:] 获取 [i+1 到 末尾] 的元素
			// ... 表示将其打散依次接进去
			tasks = append(tasks[:i], tasks[i+1:]...)
			recordLog(fmt.Sprintf("删除任务 [%s]", id))
			fmt.Println("任务删除成功！")
			return
		}
	}
	fmt.Println("错误：未找到指定的任务ID。")
}

// main 主函数，整个程序的唯一入口
func main() {
	// 第一步：先将硬盘的文件数据反序列化读取到内存里
	loadData()

	// for { ... } 是Go语言里的死循环（等同于 while (true)），用于提供常驻菜单
	for {
		fmt.Println("\n========== 任务管理系统 ==========")
		fmt.Println("1. 新增任务")
		fmt.Println("2. 查看任务")
		fmt.Println("3. 更新状态")
		fmt.Println("4. 删除任务")
		fmt.Println("5. 校验底层操作记录防篡改链")
		fmt.Println("0. 退出系统")
		fmt.Println("==================================")

		choice := readInput("请选择操作编号: ")

		// switch 分支语句，根据用户输入的字符串数字执行对应的函数
		switch choice {
		case "1":
			addTask()
		case "2":
			queryTasks()
		case "3":
			updateTask()
		case "4":
			deleteTask()
		case "5":
			// 重新加载数据，以便检测在程序运行期间文件是否被外部记事本等篡改
			loadData()
			if verifyChain() {
				fmt.Println("验证结果：日志链完整，所有历史数据均未被外部篡改！")
			} else {
				fmt.Println("警告检测：发现链条断裂或指纹不匹配，记录文件可能已被文本编辑器强行修改过！")
			}
		case "0":
			fmt.Println("感谢使用，系统已退出。")
			return // 终止整个 main 函数，程序结束
		default:
			fmt.Println("无效的输入编号。")
		}
	}
}
