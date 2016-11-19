package config

type MailGunConfig struct {
	Domain        string
	PrivateApiKey string
	PublicApiKey  string
}

func Config() MailGunConfig {
	config := MailGunConfig{
		"domain",
		"privkey",
		"pubkey",
	}
	return config
}
