package helpers

import (
	"database/sql"
	"os"
)

type Todo struct {
	Name string
	ID int
}

var dbLoc = os.Getenv("TODO_DATABASE_LOCATION")

type App struct {
	DB *sql.DB
}

var Application = App{}

var schema = `
CREATE TABLE IF NOT EXISTS todo (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, name varchar(255));
`

