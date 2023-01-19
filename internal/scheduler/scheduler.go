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
	ticker := time.NewTicker(time.Duration(5) * time.Second)
	for {
		select {
		case command := <-appCore.Command:
			result := ""
			//todo handle error
			service, _ := projectSeletor(appCore.Cfg, command.ChatID)
			if service != nil {
				if command.Type == "button_check_availability" {
					result, _ = availability.CheckAvailabilityEnv(service, false)
				}
				if command.Type == "button_check_ssl" {
					result, _ = sslchecker.CheckSSLEnv(service, false)
				}
				if command.Type == "button_check_billing" {
					result, _ = selectel.CheckBillingMessage(service, false)
				}
				appCore.Message <- app.Message{
					Type:    "response",
					Payload: result,
					ChatID:  command.ChatID,
				}
			}

		case <-ticker.C:
			for key, _ := range appCore.Cfg.Services {
				result, doSendReport := availability.CheckAvailabilityEnv(&appCore.Cfg.Services[key], true)
				if doSendReport {
					appCore.Message <- app.Message{
						Type:    "response",
						Payload: result,
						ChatID:  appCore.Cfg.Services[key].TelegramChatId,
					}
				}

				result, doSendReport = sslchecker.CheckSSLEnv(&appCore.Cfg.Services[key], true)
				if doSendReport {
					appCore.Message <- app.Message{
						Type:    "response",
						Payload: result,
						ChatID:  appCore.Cfg.Services[key].TelegramChatId,
					}
				}

				result, doSendReport = selectel.CheckBillingMessage(&appCore.Cfg.Services[key], true)
				if doSendReport {
					appCore.Message <- app.Message{
						Type:    "response",
						Payload: result,
						ChatID:  appCore.Cfg.Services[key].TelegramChatId,
					}
				}

				// todo check ssl and billing
			}
		}
	}
}

func projectSeletor(services *config.Config, chatID int) (*config.Services, error) {
	for key, service := range services.Services {
		if service.TelegramChatId == chatID {
			return &services.Services[key], nil
		}
	}
	return nil, errors.New("Chat for project not found")
}
