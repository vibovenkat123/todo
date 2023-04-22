package helpers

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func (app *App) Initialize() {
	err := app.connectToDB()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	err = app.execSchema(schema)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}

func (app *App) connectToDB() error {
	if len(dbLoc) == 0 {
		return errors.New("Set the environment variable `TODO_DATABASE_LOCATION`")
	}
	db, err := sql.Open("sqlite3", dbLoc)
	if err != nil {
		return fmt.Errorf("Error while opening file %s: %s", dbLoc, err)
	}
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Couldn't ping database: %s", err)
	}
	app.DB = db
	return nil
}

func (app *App) execSchema(schema string) error {
	if app.DB == nil {
		return errors.New("No db")
	}
	_, err := app.DB.Exec(schema)
	if err != nil {
		return fmt.Errorf("Error executing schema: %s", err)
	}
	return nil
}

func (app *App) GetAllTodos() ([]Todo, error) {
	if app.DB == nil {
		return nil, errors.New("No db")
	}
	todos := []Todo{}
	rows, err := app.DB.Query("SELECT id, name FROM TODO")
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return nil, fmt.Errorf("Error while scanning the `todo` table: %s", err)
		}
		todo := Todo{
			Name: name,
			ID:   id,
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (app *App) InsertTodo(name string) (error) {
	if app.DB == nil {
		return errors.New("No db")
	}
	stmt, err := app.DB.Prepare("INSERT INTO todo (name) VALUES (?)")
	if err != nil {
		return fmt.Errorf("Error preparing statment: %s", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(name)
	if err != nil {
		return fmt.Errorf("Error inserting into table: %s", err)
	}
	return nil
}
