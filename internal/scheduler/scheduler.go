package scheduler

import (
	"context"
	"github.com/px2x/infrastructure-tg-notifier/internal/app"
	"github.com/px2x/infrastructure-tg-notifier/internal/availability"
	"time"
)

func Run(ctx *context.Context, app2 *app.App) {
	ticker := time.NewTicker(time.Duration(1) * time.Minute)
	for {
		select {
		case command := <-app2.Command:
			if command.Type == "button_check_availability" {
				result := availability.CheckAvailability(app2.Cfg.Services[0].Env[0].Link[0].Url)
				println(result)

				app2.Message <- app.Message{
					Type:    "response",
					Payload: "Checked" + app2.Cfg.Services[0].Env[0].Link[0].Url,
					ChatID:  command.ChatID,
				}
			}
		case <-ticker.C:
			// todo
		}
	}
}
