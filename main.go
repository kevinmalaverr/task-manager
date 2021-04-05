package main

import "TaskManager/console"

func main() {
	t1 := &Task{
		name:        "example",
		description: "example description",
		completed:   false,
	}

	t2 := &Task{
		name:        "example 2",
		description: "example description 2",
		completed:   true,
	}

	t3 := &Task{
		name:        "third",
		description: "example description 2",
		completed:   true,
	}

	tasks := &TaskList{
		list: []*Task{
			t1, t2,
		},
	}

	tasksMap := make(map[string]*TaskList)

	tasksMap["kevin"] = tasks

	tasksMap["kevin"].addToList(t3)

	tasksMap["kevin"].printCompletedList()

	options := []string{"a", "b"}

	res, err := console.Select("options", options)

	if err != nil {
		print(err.Error())
	}

	println(res)
}
