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
	"github.com/zuzuleinen/dave/config"
)

type ExactTime struct {
	hours   int
	minutes int
}

func Remind(db *sql.DB) {
	fmt.Print("Task name: ")

	reader := bufio.NewReader(os.Stdin)
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSuffix(taskName, "\n")


	exactTime := scanExactTime()
	t := time.Date(scanYear(), scanMonth(), scanDay(), exactTime.hours, exactTime.minutes, 0, 0, time.UTC)

	dbItems := []reminder.Reminder{{taskName, t.Format(config.TimeFormat())}}

	reminder.Save(db, dbItems)
	color.Green("Great I will remind you to `%s`. Time: %s.", taskName, t.Format(config.TimeFormat()))
}

func scanYear() int {
	var y int
	var prompt string

	fmt.Print(fmt.Sprintf("Year[%d]:", time.Now().Year()))
	fmt.Scanln(&prompt)

	if prompt == "" {
		y = time.Now().Year()
	} else {
		y, _ = strconv.Atoi(prompt)
	}

	return y
}

func scanMonth() time.Month {
	var prompt string
	var m time.Month

	fmt.Print(fmt.Sprintf("Month[%d]:", time.Now().Month()))
	fmt.Scanln(&prompt)

	if prompt == "" {
		m = time.Now().Month()
	} else {
		monthIndex, _ := strconv.Atoi(prompt)
		m = time.Month(monthIndex)
	}

	return m
}

func scanExactTime() ExactTime {
	var prompt string

	fmt.Print("Time(HH:mm):")
	fmt.Scanln(&prompt)

	res := strings.Split(prompt, ":")

	h, _ := strconv.Atoi(res[0])
	m, _ := strconv.Atoi(res[1])

	return ExactTime{hours: h, minutes: m}
}

func scanDay() int {
	var d int
	var prompt string

	fmt.Print(fmt.Sprintf("Day[%d]:", time.Now().Day()))
	fmt.Scanln(&prompt)

	if prompt == "" {
		d = time.Now().Day()
	} else {
		d, _ = strconv.Atoi(prompt)
	}

	return d
}
