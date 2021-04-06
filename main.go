package main

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// tasks := &TaskList{}
	// menu := Menu{}

	MakeMigrations()

	db := GetConnection()
	q := `INSERT INTO notes (title, description, updated_at)
            VALUES(?, ?, ?)`
	// Preparamos la petición para insertar los datos de manera
	// segura
	stmt, err := db.Prepare(q)
	if err != nil {
		return
	}
	// Nos aseguramos de cerrar el recurso antes de finalizar la
	// función gracias a defer.
	defer stmt.Close()
	r, err := stmt.Exec("title", "des", time.Now())
	if err != nil {
		return
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return
	}

	// finished := false

	// for !finished {
	// 	finished = menu.intro(tasks)
	// }

}

// tasksMap := make(map[string]*TaskList)

// tasksMap["kevin"] = tasks

// tasksMap["kevin"].addToList(t3)

// tasksMap["kevin"].printCompletedList()
