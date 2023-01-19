package tgbot

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"math/rand"
	"strconv"
	"time"
)

func (passer *DataPasser) defaultHandler(ctx context.Context, tgBot *bot.Bot, update *models.Update) {
	return
}

func (passer *DataPasser) helpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "Используйте команды \n" +
			"/start - проверить данные по проекту \n " +
			"/getid - получить ID этого чата",
	})
}

func (passer *DataPasser) getIdHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "ID этого чата \n" + strconv.Itoa(update.Message.Chat.ID),
	})
}

func (passer *DataPasser) luckHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	reasons := []string{
		"🏀",
		"🎯",
		"🎳",
		"⚽",
	}
	time.Sleep(time.Duration(rand.Intn(50)+7) * time.Second)
	b.SendDice(ctx, &bot.SendDiceParams{
		ChatID:           update.Message.Chat.ID,
		ReplyToMessageID: update.Message.ID,
		Emoji:            reasons[rand.Intn(len(reasons))],
	})
}

func (passer *DataPasser) checkHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	cb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Биллинг хостера", CallbackData: "button_check_billing"},
				{Text: "SSL серт", CallbackData: "button_check_ssl"},
				{Text: "Доступность", CallbackData: "button_check_availability"},
			},
		},
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Что проверим?",
		ReplyMarkup: cb,
	})
}
