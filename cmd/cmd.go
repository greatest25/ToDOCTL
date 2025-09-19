package cmd

import (
	"fmt"
	"time"

	"github.com/greatest25/ToDoCTL/todo"
)

// TodoManager 管理待办事项列表
type TodoManager struct {
	Todos []todo.Item
}

// NewTodoManager 创建新的TodoManager
func NewTodoManager() *TodoManager {
	return &TodoManager{
		Todos: []todo.Item{},
	}
}

// AddTodo 添加新的待办事项
func (m *TodoManager) AddTodo(task string) {
	newTodo := todo.Item{
		ID:          len(m.Todos) + 1,
		Task:        task,
		Completed:   false,
		CreatedDate: time.Now(),
	}
	m.Todos = append(m.Todos, newTodo)
}

func (m *TodoManager) ListTodos() {
	if len(m.Todos) == 0 {
		fmt.Println("没有待办事项")
		return
	}
	fmt.Println("待办事项列表:")
	for i, item := range m.Todos {
		status := "[ ]"
		if item.Completed {
			status = "[✓]"
		}
		fmt.Printf("%d. %s %s (创建于: %s)\n",
			i+1, status, item.Task, item.CreatedDate.Format("2006-01-02 15:04:05"))
	}
}

func (m *TodoManager) CompletedTodos(id int) error {
	if id < 0 || id >= len(m.Todos) {
		return fmt.Errorf("无效的ID: %d", id)
	}
	m.Todos[id].Completed = true
	m.Todos[id].CompletedDate = time.Now()
	return nil
}

func (m *TodoManager) DeleteTodo(id int) error {
	if id < 0 || id >= len(m.Todos) {
		return fmt.Errorf("无效的ID: %d", id)
	}
	m.Todos = append(m.Todos[:id], m.Todos[id+1:]...)
	return nil
}
