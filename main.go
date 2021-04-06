package main

func main() {
	// t1 := &Task{
	// 	name:        "example",
	// 	description: "example description",
	// 	completed:   false,
	// }

	// t2 := &Task{
	// 	name:        "example 2",
	// 	description: "example description 2",
	// 	completed:   true,
	// }

	// t3 := &Task{
	// 	name:        "third",
	// 	description: "example description 2",
	// 	completed:   true,
	// }

	tasks := &TaskList{}

	// tasksMap := make(map[string]*TaskList)

	// tasksMap["kevin"] = tasks

	// tasksMap["kevin"].addToList(t3)

	// tasksMap["kevin"].printCompletedList()

	menu := Menu{}

	finished := false

	for !finished {
		finished = menu.intro(tasks)
	}

}
