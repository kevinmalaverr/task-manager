package main

import (
	"TaskManager/lib"
	"TaskManager/services"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	tasks := &services.TaskList{}
	menu := Menu{}

	lib.MakeMigrations()

	finished := false

	for !finished {
		finished = menu.intro(tasks)
	}

}

// tasksMap := make(map[string]*TaskList)

// tasksMap["kevin"] = tasks

// tasksMap["kevin"].addToList(t3)

// tasksMap["kevin"].printCompletedList()
