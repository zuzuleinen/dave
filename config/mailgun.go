package config

type MailGunConfig struct {
	Domain        string
	PrivateApiKey string
	PublicApiKey  string
}

func Config() MailGunConfig {
	config := MailGunConfig{
		"sandbox0071112e3f864716b83cd8ae6de50cd9.mailgun.org",
		"key-da9d400c2fde42a612fdc6e0b4612017",
		"pubkey-b272ce06375dc99cd552c08eec73f22f",
	}
	return config
}
