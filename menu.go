package main

import (
	"TaskManager/console"
)

type Menu struct{}

func (menu *Menu) intro(tasks *TaskList) bool {
	options := []string{"Add task", "Show all task", "Exit"}

	console.PrintHeader("Task Manager")
	res, _ := console.Select("select an option", options)

	switch res {
	case options[0]:
		menu.addTask(tasks)
	case options[1]:
		tasks.printList()
		console.Input("enter to go back")
	default:
		print("See you later!")
		return true
	}
	return false
}

func (menu *Menu) addTask(tasks *TaskList) {
	task := &Task{}

	task.name, _ = console.Input("Name: ")
	task.description, _ = console.Input("Description: ")
	task.completed = false

	task.print()

	res, _ := console.Select("Confirm?", []string{"yes", "no"})

	if res == "yes" {
		tasks.addToList(task)
	}
}
