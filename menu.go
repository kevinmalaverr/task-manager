package main

import (
	"TaskManager/console"
)

type Menu struct{}

func (menu *Menu) intro(tasks *TaskList) bool {
	options := []string{"Add task", "Show tasks", "Exit"}

	console.PrintHeader("Task Manager")
	res, _ := console.Select("select an option", options)

	switch res {
	case options[0]:
		menu.addTask(tasks)
	case options[1]:
		menu.showTasks(tasks)
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

func (menu *Menu) showTasks(tasks *TaskList) {
	options := []string{"All", "Completed", "Uncompleted"}
	res, _ := console.Select("Which tasks do you want to see?", options)

	switch res {
	case options[0]:
		tasks.printList()
	case options[1]:
		tasks.printCompletedList()
	case options[1]:
		tasks.printUncompletedList()
	default:
		return
	}

	console.Continue()
}
