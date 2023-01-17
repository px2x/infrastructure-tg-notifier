package scheduler

import (
	"context"
	"errors"
	"github.com/px2x/infrastructure-tg-notifier/config"
	"github.com/px2x/infrastructure-tg-notifier/internal/app"
	"github.com/px2x/infrastructure-tg-notifier/internal/availability"
	"github.com/px2x/infrastructure-tg-notifier/internal/selectel"
	"github.com/px2x/infrastructure-tg-notifier/internal/sslchecker"
	"time"
)

func Run(ctx *context.Context, appCore *app.App) {
	ticker := time.NewTicker(time.Duration(1) * time.Minute)
	for {
		select {
		case command := <-appCore.Command:
			result := ""
			//todo handle error
			service, _ := projectSeletor(appCore.Cfg.Services, command.ChatID)
			if command.Type == "button_check_availability" {

				result = availability.CheckAvailabilityEnv(service)
			}
			if command.Type == "button_check_ssl" {
				result = sslchecker.CheckSSLEnv(service)
			}
			if command.Type == "button_check_billing" {
				result = selectel.CheckBillingMessage(service.SelectelAPIKey)
			}
			appCore.Message <- app.Message{
				Type:    "response",
				Payload: result,
				ChatID:  command.ChatID,
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
