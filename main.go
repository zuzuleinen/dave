package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/zuzuleinen/dave/database"
	"github.com/zuzuleinen/dave/prompt"
)

const DB_PATH = "davedb.db"

func main() {
	db := database.Connect(DB_PATH)
	defer db.Close()

	prompt.SayHello()
	command := prompt.AskForCommand()
	prompt.ObeyCommand(command, db)
}


