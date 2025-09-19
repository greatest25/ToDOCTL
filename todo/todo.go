// filepath: f:\Download\golearning\ToDoCTL\todo\todo.go
package todo

import "time"

// Item 表示一个待办事项
type Item struct {
	ID            int
	Task          string
	Completed     bool
	CreatedDate   time.Time
	CompletedDate time.Time // 添加完成日期字段
}
