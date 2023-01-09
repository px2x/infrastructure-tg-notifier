package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	tgBot, err := bot.New("5858864198:AAHNKHhpXC2yZ9ZPHhjjnBlRk-y27LGoQ5Y", opts...)
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
