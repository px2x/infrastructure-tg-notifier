package config

import "time"

type App struct {
	TelegramAPIKey               string        `yaml:"telegramAPIKey"`
	CheckIntervalSSL             time.Duration `yaml:"checkIntervalSSL"`
	CheckIntervalSelectelBilling time.Duration `yaml:"checkIntervalSelectelBilling"`
	CheckIntervalSiteDefault     time.Duration `yaml:"checkIntervalSiteDefault"`
}
