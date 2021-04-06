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
	if len(tl.list) == 0 {
		fmt.Println("	⚠ the list is empty")
		fmt.Println()
		return
	}
	for _, task := range tl.list {
		task.print()
	}
}

func (tl *TaskList) printCompletedList() {
	console.PrintHeader("Completed Tasks")
	if len(tl.list) == 0 {
		fmt.Println("	⚠ the list is empty")
		fmt.Println()
		return
	}
	for _, task := range tl.list {
		if task.completed {
			task.print()
		}
	}
}

func (tl *TaskList) printUncompletedList() {
	console.PrintHeader("Uncompleted Tasks")
	if len(tl.list) == 0 {
		fmt.Println("	⚠ the list is empty")
		fmt.Println()
		return
	}
	for _, task := range tl.list {
		if !task.completed {
			task.print()
		}
	}
}
