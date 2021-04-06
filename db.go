package main

import "database/sql"

// Puntero a la estructura DB, nos permite manejar la
// base de datos
var db *sql.DB

func GetConnection() *sql.DB {
	// Para evitar realizar una nueva conexión en cada llamada a
	// la función GetConnection.
	if db != nil {
		return db
	}
	// Declaramos la variable err para poder usar el operador
	// de asignación “=” en lugar que el de asignación corta,
	// para evitar que cree una nueva variable db en este scope y
	// en su lugar que inicialice la variable db que declaramos a
	// nivel de paquete.
	var err error
	// Conexión a la base de datos
	db, err = sql.Open("sqlite3", "data.sqlite")
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
					created_at TIMESTAMP DEFAULT DATETIME,
					updated_at TIMESTAMP NOT NULL
			 );`
	_, err := db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}
