package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/px2x/infrastructure-tg-notifier/config"
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

	println(cfg)

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	tgBot, err := bot.New("", opts...)
	if err != nil {
		panic(err)
	}

	tgBot.Start(ctx)
}

func handler(ctx context.Context, tgBot *bot.Bot, update *models.Update) {
	tgBot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
