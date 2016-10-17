package commands

import (
	"github.com/zuzuleinen/dave/reminder"
	"fmt"
	"database/sql"
	"github.com/zuzuleinen/dave/credentials"
)

func InstallDatabase(db *sql.DB) {
	reminder.CreateTable(db)
	fmt.Println("Created reminders table.")
	credentials.CreateTable(db)
	fmt.Println("Created credentials table.")
}
