package config

import "time"

type Env struct {
	Name string `yaml:"name"`
	Link []Link `yaml:"links"`
}

type Link struct {
	Url               string        `yaml:"url"`
	CheckIntervalSSL  time.Duration `yaml:"checkIntervalSSL"`
	CheckIntervalSite time.Duration `yaml:"checkIntervalSite"`
}

type Services struct {
	Name           string `yaml:"name"`
	TelegramChatId string `yaml:"telegramChatId"`
	SelectelAPIKey string `yaml:"selectelAPIKey"`
	Env            []Env  `yaml:"envs"`
}
