package services

import (
	"TaskManager/lib"
	"errors"
	"fmt"
)

type Task struct {
	Id          int
	Name        string
	Description string
	Completed   bool
}

func (t *Task) Complete() {
	t.Completed = true
}

func (t *Task) UpdateDescription(newDescription string) {
	t.Description = newDescription
}

func (t *Task) UpdateName(newName string) {
	t.Name = newName
}

func (self *Task) Print() {
	fmt.Println("Name:", self.Name)
	fmt.Println("Description:", self.Description)
	fmt.Println("Completed:", self.Completed)
	fmt.Println()
}

func (t *Task) Insert() error {
	db := lib.GetConnection()

	q := `INSERT INTO tasks (name, description, completed)
            VALUES(?, ?, ?)`

	stmt, err := db.Prepare(q)
	if err != nil {
		print(err.Error())
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(t.Name, t.Description, t.Completed)
	if err != nil {
		return err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: expected an affected row")
	}
	return nil
}
