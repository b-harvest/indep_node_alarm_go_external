package alert

import (
	"indep_node_alarm_go_external/pkg/config"

	"github.com/rs/zerolog/log"

	"github.com/PagerDuty/go-pagerduty"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendAlerts(err error, config config.Config, server_name string, address string) {
	// send alert via Telegram
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to create Telegram bot")
	}
	msg_text := "ALERT: Unable to connect to server " + server_name + " at " + address
	msg := tgbotapi.NewMessage(config.ChatID, msg_text)
	bot.Send(msg)

	// send alert via PagerDuty
	client := pagerduty.NewClient(config.PagerDutyToken)
	incident := pagerduty.CreateIncidentOptions{
		Type:  "incident",
		Title: "ALERT: Unable to connect to server",
	}
	if _, err := client.CreateIncident(config.PagerDutyUserID, &incident); err != nil {
		log.Error().Err(err).Msg("Failed to create PagerDuty incident")
		return
	}
}
