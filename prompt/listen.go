package prompt

import (
	"database/sql"
	"github.com/fatih/color"
	"github.com/zuzuleinen/dave/prompt/commands"
	"os"
	"github.com/zuzuleinen/dave/reminder"
	"github.com/zuzuleinen/dave/config"
	"fmt"
)

func ObeyCommand(command string, db *sql.DB) {
	switch command {
	case "install":
		commands.InstallDatabase(db)
	case "remind":
		commands.Remind(db)
	case "reminders":
		commands.ListReminders(db)
	case "credential":
		commands.Credential(db)
	case "credentials":
		commands.ListCredentials(db)
	case "cli":
		reminder.Remind(db)
	case "--help", "-h":
		commands.Help()
	case "--version", "-v":
		fmt.Println(config.Version())
	default:
		color.Red("Sorry, I dont understand '%s' command.", command)
		commands.Help()
		os.Exit(1)
	}
}
