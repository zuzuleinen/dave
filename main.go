package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/zuzuleinen/dave/database"
	"github.com/zuzuleinen/dave/prompt"
	"os"
	"github.com/zuzuleinen/dave/reminder"
	"github.com/zuzuleinen/dave/config"
)

func main() {
	db := database.Connect(config.DbPath())
	defer db.Close()

	if contains(os.Args, "cli") {
		reminder.Remind(db)
	} else {
		prompt.SayHello()
		command := prompt.AskForCommand()
		prompt.ObeyCommand(command, db)
	}
}

func contains(s []string, e string) bool {
	for _, value := range s {
		if e == value {
			return true
		}
	}
	return false
}
