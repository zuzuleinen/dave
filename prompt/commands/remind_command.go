package commands

import (
	"github.com/zuzuleinen/dave/reminder"
	"strings"
	"os"
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"database/sql"
)

func Remind(db *sql.DB) {
	fmt.Print("Task name: ")
	reader := bufio.NewReader(os.Stdin)
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSuffix(taskName, "\n")
	dbItems := []reminder.Reminder{{taskName}}

	reminder.Save(db, dbItems)
	color.Green("Great I will remind you to %s.", taskName)
}
