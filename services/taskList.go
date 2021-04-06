package services

import (
	"TaskManager/lib"
	console "TaskManager/lib"
	"fmt"
)

type TaskList struct {
	list []*Task
}

func (tl *TaskList) AddToList(t *Task) {
	tl.list = append(tl.list, t)
}

func (tl *TaskList) RemoveFromList(index int) {
	tl.list = append(tl.list[:index], tl.list[index+1:]...)
}

func (tl *TaskList) PrintList() {
	console.PrintHeader("All Tasks")
	if len(tl.list) == 0 {
		fmt.Println("	⚠ the list is empty")
		fmt.Println()
		return
	}
	for _, task := range tl.list {
		task.Print()
	}
}

func (tl *TaskList) PrintCompletedList() {
	tl.GetAll()
	console.PrintHeader("Completed Tasks")
	if len(tl.list) == 0 {
		fmt.Println("	⚠ the list is empty")
		fmt.Println()
		return
	}
	for _, task := range tl.list {
		if task.Completed {
			task.Print()
		}
	}
}

func (tl *TaskList) PrintUncompletedList() {
	console.PrintHeader("Uncompleted Tasks")
	if len(tl.list) == 0 {
		fmt.Println("	⚠ the list is empty")
		fmt.Println()
		return
	}
	for _, task := range tl.list {
		if !task.Completed {
			task.Print()
		}
	}
}

func (tl *TaskList) GetAll() error {
	db := lib.GetConnection()
	q := `SELECT
            id, name, description, completed
            FROM tasks`

	rows, err := db.Query(q)
	if err != nil {
		print(err.Error())
		return err
	}

	defer rows.Close()

	task := Task{}
	tl.list = []*Task{}

	for rows.Next() {
		rows.Scan(
			&task.Id,
			&task.Name,
			&task.Description,
			&task.Completed,
		)

		tl.list = append(tl.list, &task)
	}

	return nil
}
