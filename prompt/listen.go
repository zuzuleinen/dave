package prompt

import (
	"fmt"
	"database/sql"
	"github.com/fatih/color"
	"github.com/zuzuleinen/dave/prompt/commands"
)

func SayHello() {
	fmt.Println("Yes, sir?")
}

func AskForCommand() string {
	var command string
	fmt.Print("\t")
	fmt.Scan(&command)
	return command
}

func ObeyCommand(command string, db *sql.DB) {
	switch command {
	case "list":
		commands.List()
	case "install":
		commands.InstallDatabase(db)
	case "remind":
		commands.Remind(db)
	case "reminders":
		commands.ListReminders(db)
	default:
		color.Red("Sorry, I dont understand '%s' command.", command)
	}
}