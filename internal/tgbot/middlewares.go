package tgbot

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
)

func (passer *DataPasser) logMessageMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			log.Printf("@%s say: %s", update.Message.From.Username, update.Message.Text)
		}
		if update.CallbackQuery != nil {
			log.Printf("@%s say: %s", update.CallbackQuery.Sender.Username, update.CallbackQuery.Data)
		}
		next(ctx, b, update)
	}
}
