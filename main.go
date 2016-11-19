package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/zuzuleinen/dave/database"
	"github.com/zuzuleinen/dave/prompt"
	"os"
	"github.com/zuzuleinen/dave/config"
)

type Command struct {
	command  string
	argument string
}

func main() {
	db := database.Connect(config.DbPath())
	defer db.Close()
	c := NewCommand()
	prompt.ObeyCommand(c.command, c.argument, db)
}

func NewCommand() *Command {
	c := Command{command:"--help"}

	if len(os.Args) > 1 {
		c.command = os.Args[1]
	}

	if len(os.Args) > 2 {
		c.argument = os.Args[2]
	}

	return &c
}