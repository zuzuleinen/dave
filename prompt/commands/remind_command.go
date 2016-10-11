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

type ExactTime struct {
	hours   int
	minutes int
}

func Remind(db *sql.DB) {
	fmt.Print("Task name: ")

	reader := bufio.NewReader(os.Stdin)
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSuffix(taskName, "\n")
	dbItems := []reminder.Reminder{{taskName}}

	year := scanYear()
	month := scanMonth()
	day := scanDay()
	exactTime := scanExactTime()

	fmt.Println(exactTime)

	reminder.Save(db, dbItems)
	color.Green("Great I will remind you to %s.", taskName)
	fmt.Println("Time is", year, day, month.String(), exactTime)
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

	hours, _ := strconv.Atoi(res[0])
	minutes, _ := strconv.Atoi(res[1])

	return ExactTime{hours, minutes}
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
