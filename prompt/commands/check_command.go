package commands

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/zuzuleinen/dave/email"
	"github.com/zuzuleinen/dave/config"
)

func Check(argument string) {
	if argument == "" {
		return;
	}
	sendEmail(argument)
	fmt.Println(color.GreenString("Email sent."))
	fmt.Println(color.GreenString(argument))

}


func sendEmail(body string) {
	subject := "check this"
	email.Send(config.YourEmail(), subject, body)
}
