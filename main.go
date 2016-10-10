package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/mitchellh/go-homedir"
	"github.com/zuzuleinen/dave/database"
	"github.com/zuzuleinen/dave/prompt"
	"os"
)

const DB_FILE = "davedb.db"

func main() {
	db := database.Connect(dbPath())
	defer db.Close()

	//todo
	//email.Send("andrey.boar@gmail.com", "Reminder", "Some plain text")

	prompt.SayHello()
	command := prompt.AskForCommand()
	prompt.ObeyCommand(command, db)
}

func dbPath() string {
	homeDir, _ := homedir.Dir()
	return homeDir + string(os.PathSeparator) + DB_FILE
}
