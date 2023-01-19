package config

import "time"

type Env struct {
	Name string `yaml:"name"`
	Link []Link `yaml:"links"`
}

type Link struct {
	Url string `yaml:"url"`
}

type Services struct {
	Name                         string        `yaml:"name"`
	TelegramChatId               int           `yaml:"telegramChatId"`
	SelectelAPIKey               string        `yaml:"selectelAPIKey"`
	CheckIntervalSSL             time.Duration `yaml:"checkIntervalSSL"`
	CheckIntervalSite            time.Duration `yaml:"checkIntervalSite"`
	CheckIntervalSelectelBilling time.Duration `yaml:"checkIntervalSelectelBilling"`
	LastCheckBilling             time.Time
	LastCheckSSL                 time.Time
	LastCheckSite                time.Time
	Env                          []Env `yaml:"envs"`
}
