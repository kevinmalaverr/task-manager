package main

type task struct {
	name        string
	description string
	completed   bool
}

type taskList struct {
	list []*task
}

func (tl *taskList) addToList(t *task) {
	tl.list = append(tl.list, t)
}

func (tl *taskList) removeFromList(index int) {
	tl.list = append(tl.list[:index], tl.list[index+1:]...)
}

func (t *task) complete() {
	t.completed = true
}

func (t *task) updateDescription(newDescription string) {
	t.description = newDescription
}

func (t *task) updateName(newName string) {
	t.name = newName
}

func main() {
	t1 := &task{
		name:        "example",
		description: "example description",
		completed:   false,
	}

	t2 := &task{
		name:        "example 2",
		description: "example description 2",
		completed:   true,
	}

	t3 := &task{
		name:        "third",
		description: "example description 2",
		completed:   true,
	}

	tasks := &taskList{
		list: []*task{
			t1, t2,
		},
	}

	tasks.addToList(t3)

	println(len(tasks.list))

}
