package main

import (
	"context"
	"github.com/px2x/infrastructure-tg-notifier/config"
	"github.com/px2x/infrastructure-tg-notifier/internal/app"
	"github.com/px2x/infrastructure-tg-notifier/internal/scheduler"
	"github.com/px2x/infrastructure-tg-notifier/internal/tgbot"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	service := app.NewApp(cfg)
	tgbot.Run(&ctx, service)
	scheduler.Run(&ctx, service)

	forever := make(chan bool)
	<-forever

}
