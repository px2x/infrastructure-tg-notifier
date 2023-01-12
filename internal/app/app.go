package app

import (
	"github.com/px2x/infrastructure-tg-notifier/config"
)

type App struct {
	Cfg     *config.Config
	Message chan Message
	Command chan Message
}

type Message struct {
	Type    string
	Payload string
	ChatID  int
}

func NewApp(cfg *config.Config) *App {
	return &App{
		Cfg:     cfg,
		Message: make(chan Message),
		Command: make(chan Message),
	}
}
