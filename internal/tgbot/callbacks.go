package tgbot

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/px2x/infrastructure-tg-notifier/internal/app"
)

func (passer *DataPasser) callbackCheckHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	passer.app.Command <- app.Message{
		Type:    update.CallbackQuery.Data,
		Payload: "",
		ChatID:  update.CallbackQuery.Message.Chat.ID,
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Chat.ID,
		Text:   "You selected the button: " + update.CallbackQuery.Data,
	})
}
