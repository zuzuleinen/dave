package commands

import (
	"database/sql"
	"fmt"
	"github.com/zuzuleinen/dave/reminder"
)

func ListReminders(db *sql.DB) {
	items := reminder.Read(db)
	for _, item := range items {
		fmt.Println(item.Name, "|", item.Time)
	}
}
