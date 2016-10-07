package config

type MailGunConfig struct {
	Domain        string
	PrivateApiKey string
	PublicApiKey  string
}

func Config() MailGunConfig {
	config := MailGunConfig{"yourDomain", "privateKey", "publicKey"}
	return config
}