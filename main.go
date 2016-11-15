package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/zuzuleinen/dave/database"
	"github.com/zuzuleinen/dave/prompt"
	"os"
	"github.com/zuzuleinen/dave/config"
)

func main() {
	db := database.Connect(config.DbPath())
	defer db.Close()

	var command string

	if (len(os.Args) > 1) {
		command = os.Args[1]
	} else {
		command = "list"
	}

	prompt.ObeyCommand(command, db)
}
