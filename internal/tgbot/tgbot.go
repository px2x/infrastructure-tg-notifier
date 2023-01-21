package tgbot

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/px2x/infrastructure-tg-notifier/internal/app"
	"regexp"
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
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, passer.checkHandler),
		bot.WithMessageTextHandler("/getid", bot.MatchTypeExact, passer.getIdHandler),
		//bot.WithMessageTextHandler("повезет ли нам?", bot.MatchTypeExact, passer.luckHandler),
		bot.WithCallbackQueryDataHandler("button_check_", bot.MatchTypePrefix, passer.callbackCheckHandler),
		bot.WithMiddlewares(passer.logMessageMiddleware),
	}

	tgBot, err := bot.New(app.Cfg.App.TelegramAPIKey, opts...)
	if err != nil {
		panic(err)
	}

	tgBot.RegisterHandlerRegexp(bot.HandlerTypeMessageText, regexp.MustCompile(`(?m)повезет ли нам`), passer.luckHandler)

	go tgBot.Start(*ctx)

	go func() {
		for message := range app.Message {
			if len(message.Payload) > 0 {
				tgBot.SendMessage(*ctx, &bot.SendMessageParams{
					ChatID:                message.ChatID,
					Text:                  message.Payload,
					ParseMode:             models.ParseModeHTML,
					DisableWebPagePreview: true,
				})
			}
		}
	}()
}
