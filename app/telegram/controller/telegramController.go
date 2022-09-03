package controller

import (
	"encoding/json"
	"main/telegram"
	"main/telegram/command"
	"main/telegram/entity"
	"net/http"
)

var bot telegram.BotInterface = telegram.Bot{}

func Webhook(w http.ResponseWriter, r *http.Request) {
	var telegramUpdate entity.TelegramUpdate

	err := json.NewDecoder(r.Body).Decode(&telegramUpdate)
	if err != nil {
		return
	}

	command.Handle(telegramUpdate)
}

func SetWebhook(w http.ResponseWriter, r *http.Request) {
	webhookUrl := r.URL.Query().Get("url")
	w.Write(bot.SetWebhook(webhookUrl))
}

func GetWebhook(w http.ResponseWriter, r *http.Request) {
	w.Write(bot.GetWebhookInfo())
}
