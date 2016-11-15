package reminder

import (
	"fmt"
	"database/sql"
	"time"
	"log"
	"github.com/zuzuleinen/dave/email"
	"github.com/zuzuleinen/dave/config"
)

func Remind(db *sql.DB) {
	for true {
		reminders := Read(db)
		for _, r := range reminders {
			if shouldRemind(r) {
				sendReminderMail(r.Name)
				Delete(r.RowId, db)
			}
		}
	}
}

func shouldRemind(r Reminder) bool {
	t, err := time.Parse(config.TimeFormat(), r.Time)

	if err != nil {
		log.Fatalln("Cannot parse time", err)
	}
	thisYear, thisMonth, today := time.Now().Date()
	y, m, d := t.Date()
	if (thisYear == y && thisMonth == m && today == d) {
		return true
	}
	return false
}

func sendReminderMail(todo string) {
	subject := fmt.Sprintf("Reminder: %s", todo)
	body := fmt.Sprintf("Hey, you need to %s", todo)

	email.Send(config.YourEmail(), subject, body)
}