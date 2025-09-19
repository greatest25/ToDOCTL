package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/greatest25/ToDoCTL/cmd"
)

func printHelp() {
	fmt.Println("\n=== 待办事项管理器 ===")
	fmt.Println("可用命令:")
	fmt.Println("  add <任务内容>     - 添加新的待办事项")
	fmt.Println("  list              - 显示所有待办事项")
	fmt.Println("  done <ID>         - 将待办事项标记为已完成")
	fmt.Println("  delete <ID>       - 删除待办事项")
	fmt.Println("  help              - 显示帮助信息")
	fmt.Println("  exit              - 退出程序")
	fmt.Println("========================")
}

func main() {
	fmt.Println("欢迎使用待办事项管理器!")
	printHelp()

	manager := cmd.NewTodoManager()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\n> ")
		scanner.Scan()
		input := scanner.Text()

		if input == "" {
			continue
		}

		parts := strings.SplitN(input, " ", 2)
		command := parts[0]

		switch command {
		case "add":
			if len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
				fmt.Println("错误: 请提供任务内容")
				continue
			}
			task := parts[1]
			manager.AddTodo(task)
			fmt.Println("✅ 成功添加新的待办事项")

		case "list":
			manager.ListTodos()

		case "done":
			if len(parts) < 2 {
				fmt.Println("错误: 请提供待办事项ID")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("错误: ID必须是数字")
				continue
			}
			// 调整索引（用户看到的ID是从1开始，但程序中是从0开始）
			err = manager.CompletedTodos(id - 1) // 正确赋值给err变量
			if err != nil {
				fmt.Println("错误:", err)
				continue
			}
			fmt.Println("✅ 已将任务标记为完成")

		case "delete":
			if len(parts) < 2 {
				fmt.Println("错误: 请提供待办事项ID")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("错误: ID必须是数字")
				continue
			}
			// 调整索引（用户看到的ID是从1开始，但程序中是从0开始）
			err = manager.DeleteTodo(id - 1) // 正确赋值给err变量
			if err != nil {
				fmt.Println("错误:", err)
				continue
			}
			fmt.Println("✅ 已删除任务")

		case "help":
			printHelp()

		case "exit":
			fmt.Println("再见!")
			return

		default:
			fmt.Printf("未知命令: %s\n", command)
			printHelp()
		}
	}
}
