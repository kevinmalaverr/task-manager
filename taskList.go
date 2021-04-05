package main

import (
	console "TaskManager/console"
	"fmt"
)

type TaskList struct {
	list []*Task
}

func (tl *TaskList) addToList(t *Task) {
	tl.list = append(tl.list, t)
}

func (tl *TaskList) removeFromList(index int) {
	tl.list = append(tl.list[:index], tl.list[index+1:]...)
}

func (tl *TaskList) printList() {
	console.PrintHeader("All Tasks")
	fmt.Println()
	for _, task := range tl.list {
		task.print()
	}
}

func (tl *TaskList) printCompletedList() {
	console.PrintHeader("Completed Tasks")
	for _, task := range tl.list {
		if task.completed {
			task.print()
		}
	}
}
