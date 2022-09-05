package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/helper"
	"main/telegram/model"
	"main/telegram/repository"
	"main/telegram/service"
	"net/http"
	"strconv"
	"strings"
)

type Bot struct{}

type mail struct {
	ChatId    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

var logService service.LogService

func getUrl() string {
	botApiKey := helper.GetEnv("TELEGRAM_BOT", "")

	if len(botApiKey) <= 0 {
		log.Fatal("Bot apiKey not found")
	}

	return helper.GetEnv("TELEGRAM_URL", "") + "/bot" + botApiKey
}

func (b Bot) SimpleSendMessage(message string, userId string) {
	url := getUrl() + "/sendMessage"
	method := "POST"

	payloadData, _ := json.Marshal(mail{ChatId: userId, Text: message, ParseMode: "html"})
	payload := strings.NewReader(string(payloadData))
	user, err := strconv.Atoi(userId)

	if err != nil {
		return
	}

	logItem := model.Log{IsBot: true, TelegramId: user, Payload: string(payloadData)}
	repository.Create(&logItem)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}

func (b Bot) SetWebhook(url string) []byte {
	requestUrl := fmt.Sprintf("%s/bot%s/setWebhook?url=%s",
		helper.GetEnv("TELEGRAM_URL", ""),
		helper.GetEnv("TELEGRAM_BOT", ""),
		url,
	)

	response, err := http.Get(requestUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil
	}

	return body
}

func (b Bot) GetWebhookInfo() []byte {
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
		log.Fatal("Cannot read a response")
	}

	return body
}
