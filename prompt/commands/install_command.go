package commands

import (
	"github.com/zuzuleinen/dave/reminder"
	"fmt"
	"database/sql"
)

func InstallDatabase(db *sql.DB) {
	reminder.CreateTable(db)
	fmt.Println("Created reminders table.")
}
