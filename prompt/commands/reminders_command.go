package commands

import (
	"fmt"
	"github.com/zuzuleinen/dave/reminder"
	"database/sql"
)

func ListReminders(db *sql.DB) {
	items := reminder.Read(db)
	for _, item := range items {
		fmt.Println("\t", item.Name)
	}
}
