package todo

import "time"

type TodoItem struct {
	ID        int
	Task      string
	Completed bool
	CreatData time.Time
}
