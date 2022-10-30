package controller

import (
	"encoding/json"
	"io"
	"main/telegram"
	"main/telegram/command"
	"main/telegram/entity"
	"main/telegram/keyboardCommand"
	"main/telegram/service"
	"net/http"
)

var bot telegram.BotInterface = telegram.Bot{}
var logService service.LogService

func Webhook(w http.ResponseWriter, r *http.Request) {
	var telegramUpdate entity.TelegramUpdate
	var telegramUpdateInline entity.TelegramUpdateInline

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &telegramUpdate)

	if err != nil {
		return
	}

	if telegramUpdate.Message.MessageId > 0 {
		command.Handle(telegramUpdate)
		go logService.LoggingFromUpdateEntity(telegramUpdate)
		return
	}

	err = json.Unmarshal(body, &telegramUpdateInline)
	if err != nil {
		return
	}

	if telegramUpdateInline.CallbackQuery.Message.MessageId > 0 {
		keyboardCommand.Handle(telegramUpdateInline)
	}
}

func SetWebhook(w http.ResponseWriter, r *http.Request) {
	webhookUrl := r.URL.Query().Get("url")
	w.Write(bot.SetWebhook(webhookUrl))
}

func GetWebhook(w http.ResponseWriter, r *http.Request) {
	w.Write(bot.GetWebhookInfo())
}
