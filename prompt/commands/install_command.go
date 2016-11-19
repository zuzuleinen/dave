package commands

import (
	"database/sql"
	"fmt"
	"github.com/zuzuleinen/dave/credentials"
	"github.com/zuzuleinen/dave/reminder"
)

func InstallDatabase(db *sql.DB) {
	reminder.CreateTable(db)
	fmt.Println("Created reminders table.")
	credentials.CreateTable(db)
	fmt.Println("Created credentials table.")
}
