package scheduler

import (
	"context"
	"errors"
	"github.com/px2x/infrastructure-tg-notifier/config"
	"github.com/px2x/infrastructure-tg-notifier/internal/app"
	"github.com/px2x/infrastructure-tg-notifier/internal/availability"
	"time"
)

func Run(ctx *context.Context, appCore *app.App) {
	ticker := time.NewTicker(time.Duration(1) * time.Minute)
	for {
		select {
		case command := <-appCore.Command:
			if command.Type == "button_check_availability" {
				//todo handle error
				service, _ := projectSeletor(appCore.Cfg.Services, command.ChatID)
				result := availability.CheckAvailabilityEnv(service)

				//result := true
				//result := availability.CheckAvailability(appCore.Cfg.Services[0].Env[0].Link[0].Url)
				println(result)
				println(service)

				appCore.Message <- app.Message{
					Type:    "response",
					Payload: "Checked" + appCore.Cfg.Services[0].Env[0].Link[0].Url,
					ChatID:  command.ChatID,
				}
			}
		case <-ticker.C:
			// todo
		}
	}
}

func projectSeletor(services []config.Services, chatID int) (*config.Services, error) {

	for _, service := range services {
		if service.TelegramChatId == chatID {
			return &service, nil
		}
	}
	return nil, errors.New("Chat for prijject not found")
}
