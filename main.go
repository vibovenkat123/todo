package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/vibovenkat123/todo/cmd/todo"
	"github.com/vibovenkat123/todo/pkg/helpers"
)

func main() {
	helpers.Application.Initialize()
	todo.Execute()
}
