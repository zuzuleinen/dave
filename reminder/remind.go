package reminder

import (
	"fmt"
	"database/sql"
	"time"
	"log"
)

func Remind(db *sql.DB) {
	reminders := Read(db)
	for _, r := range reminders {
		if shouldRemember(r) {
			fmt.Println("Remind me to", r.Name)
		}
	}
}

func shouldRemember(r Reminder) bool {
	t, err := time.Parse("Monday, 2 Jan 2006 at 15:04", r.Time)

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