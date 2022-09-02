package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/entity"
	"main/helper"
	"net/http"
)

func Webhook(w http.ResponseWriter, r *http.Request) {
	var telegramUpdate entity.TelegramUpdate

	err := json.NewDecoder(r.Body).Decode(&telegramUpdate)
	if err != nil {
		return
	}
}

func SetWebhook(w http.ResponseWriter, r *http.Request) {
	webhookUrl := r.URL.Query().Get("url")

	url := fmt.Sprintf("%s/bot%s/setWebhook?url=%s",
		helper.GetEnv("TELEGRAM_URL", ""),
		helper.GetEnv("TELEGRAM_BOT", ""),
		webhookUrl,
	)

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	w.Write(body)
}

func GetWebhook(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/bot%s/getWebhookInfo",
		helper.GetEnv("TELEGRAM_URL", ""),
		helper.GetEnv("TELEGRAM_BOT", ""),
	)

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	w.Write(body)
}
