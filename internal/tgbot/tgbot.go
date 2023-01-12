package tgbot

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/px2x/infrastructure-tg-notifier/internal/app"
)

type DataPasser struct {
	app *app.App
}

func Run(ctx *context.Context, app *app.App) {
	passer := &DataPasser{
		app: app,
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(passer.defaultHandler),
		bot.WithMessageTextHandler("/help", bot.MatchTypeExact, passer.helpHandler),
		bot.WithMessageTextHandler("/check", bot.MatchTypeExact, passer.checkHandler),
		bot.WithMessageTextHandler("/getid", bot.MatchTypeExact, passer.getIdHandler),
		bot.WithCallbackQueryDataHandler("button_check_", bot.MatchTypePrefix, passer.callbackCheckHandler),
		bot.WithMiddlewares(passer.logMessageMiddleware),
	}

	tgBot, err := bot.New(app.Cfg.App.TelegramAPIKey, opts...)
	if err != nil {
		panic(err)
	}

	//tgBot.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, helpHandler)

	go tgBot.Start(*ctx)

	go func() {
		for message := range app.Message {
			if len(message.Payload) > 0 {
				tgBot.SendMessage(*ctx, &bot.SendMessageParams{
					ChatID: message.ChatID,
					Text:   message.Payload,
				})
			}
		}
	}()
}
