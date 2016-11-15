package commands

import "fmt"

func Help() {
	doc := `
Usage:
    dave <command>

List of commands:
  install:        Install the SQLite database.
  remind:         Create a new reminder.
  reminders:      List all pending reminders.
  credential:     Add a new credential record.
  credentials:    List all credentials.


Options:
  -h --help         Show this screen.
  -v, --version     Show version.
`

	fmt.Println(doc)
}
