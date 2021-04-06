package main

import (
	console "TaskManager/lib"
	"TaskManager/services"
)

type Menu struct{}

func (menu *Menu) intro(tasks *services.TaskList) bool {
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

func (menu *Menu) addTask(tasks *services.TaskList) {
	task := &services.Task{}

	task.Name, _ = console.Input("Name: ")
	task.Description, _ = console.Input("Description: ")
	task.Completed = false

	task.Print()

	res, _ := console.Select("Confirm?", []string{"yes", "no"})

	if res == "yes" {
		task.Insert()
	}
}

func (menu *Menu) showTasks(tasks *services.TaskList) {
	options := []string{"All", "Completed", "Uncompleted"}
	res, _ := console.Select("Which tasks do you want to see?", options)

	tasks.GetAll()

	switch res {
	case options[0]:
		tasks.PrintList()
	case options[1]:
		tasks.PrintCompletedList()
	case options[1]:
		tasks.PrintUncompletedList()
	default:
		return
	}

	console.Continue()
}
