package cmd

import {"flag"
		"./todo"}

func addTools(task string){
	newTodo:=TodoItem{
		ID:	len(m.TOdos)+1,
		Task: task,
		Completed:	false,
		CreatedData: time.Now()
	}
	m.todos=append(m.todos.newTodo)
}