package email

import (
	"github.com/mailgun/mailgun-go"
	"log"
	"github.com/zuzuleinen/dave/config"
)


func Send(from, subject, text string)  {
	mailGunConfig := config.Config()

	mg := mailgun.NewMailgun(
		mailGunConfig.Domain,
		mailGunConfig.PrivateApiKey,
		mailGunConfig.PublicApiKey,
	)

	m := mg.NewMessage(
		from,
		subject,
		text,
		config.YourEmail(),
	)

	_, _, err := mg.Send(m)

	if err != nil {
		log.Fatal(err)
	}
}
