package main

import "fmt"

type Task struct {
	name        string
	description string
	completed   bool
}

func (t *Task) complete() {
	t.completed = true
}

func (t *Task) updateDescription(newDescription string) {
	t.description = newDescription
}

func (t *Task) updateName(newName string) {
	t.name = newName
}
func (self *Task) print() {
	fmt.Println("Name:", self.name)
	fmt.Println("Descroption:", self.description)
	fmt.Println("Completed:", self.completed)
	fmt.Println()
}
