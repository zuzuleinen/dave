package prompt

import (
	"database/sql"
	"github.com/fatih/color"
	"github.com/zuzuleinen/dave/prompt/commands"
	"os"
	"github.com/zuzuleinen/dave/reminder"
)

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
	case "credential":
		commands.Credential(db)
	case "credentials":
		commands.ListCredentials(db)
	case "cli":
		reminder.Remind(db)
	default:
		color.Red("Sorry, I dont understand '%s' command.", command)
		commands.List()
		os.Exit(1)
	}
}
