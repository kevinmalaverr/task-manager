package lib

import (
	"database/sql"
)

var db *sql.DB

func GetConnection() *sql.DB {

	if db != nil {
		return db
	}

	var err error
	db, err = sql.Open("sqlite3", "data.db")
	if err != nil {
		panic(err)
	}
	return db
}

func MakeMigrations() error {
	db := GetConnection()
	q := `CREATE TABLE IF NOT EXISTS tasks (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					name VARCHAR(64) NULL,
					description VARCHAR(200) NULL,
					completed INTEGER NULL,
					created_at TIMESTAMP DEFAULT DATETIME
			 );`
	_, err := db.Exec(q)
	if err != nil {
		print(err.Error())
		return err
	}
	return nil
}
