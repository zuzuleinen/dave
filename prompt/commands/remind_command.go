package commands

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/fatih/color"
	"github.com/zuzuleinen/dave/reminder"
	"os"
	"strconv"
	"strings"
	"time"
)

func Remind(db *sql.DB) {
	fmt.Print("Task name: ")

	reader := bufio.NewReader(os.Stdin)
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSuffix(taskName, "\n")
	dbItems := []reminder.Reminder{{taskName}}

	year := scanYear()
	month := scanMonth()
	hour := scanHour()

	reminder.Save(db, dbItems)
	color.Green("Great I will remind you to %s.", taskName)
	fmt.Println("Time is", year, month, hour)
}

func scanYear() int {
	var y int
	var prompt string

	fmt.Print("Year[2016]:")
	fmt.Scanln(&prompt)

	if prompt == "" {
		y = time.Now().Year()
	} else {
		y, _ = strconv.Atoi(prompt)
	}

	return y
}

func scanMonth() string {
	var prompt string
	var m string

	fmt.Print("Month[10]:")
	fmt.Scanln(&prompt)

	if prompt == "" {
		m = time.Now().Month().String()
	} else {
		monthIndex, _ := strconv.Atoi(prompt)
		m = time.Month(monthIndex).String()
	}

	return m
}

func scanHour() string {
	var hour string

	fmt.Print("Time(HH:mm):")
	fmt.Scanln(&hour)

	return hour
}
