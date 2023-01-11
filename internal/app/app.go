package app

import (
	"github.com/px2x/infrastructure-tg-notifier/config"
)

type App struct {
	Cfg      *config.Config
	Message  chan string
	Handlers interface{}
}

func NewApp(cfg *config.Config) *App {
	return &App{
		Cfg:     cfg,
		Message: make(chan string, 5),
	}
}
