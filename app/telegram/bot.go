package telegram

import (
	"encoding/json"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"io"
	"io/ioutil"
	"log"
	"main/helper"
	"main/telegram/model"
	"main/telegram/repository"
	"main/telegram/service"
	"net/http"
	"strconv"
)

type Bot struct{}

type mail struct {
	ChatId    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
	Keyboard  string `json:"reply_markup"`
}

type photoMail struct {
	ChatId string `json:"chat_id"`
	Photo  string `json:"text"`
}

var logService service.LogService

var contentTypeMap = map[string]string{
	"image/jpeg": "photo",
	"image/png":  "photo",
	"video/mp4":  "video",
	"video/gif":  "animation",
}

func getUrl() string {
	botApiKey := helper.GetEnv("TELEGRAM_BOT", "")

	if len(botApiKey) <= 0 {
		log.Fatal("Bot apiKey not found")
	}

	return helper.GetEnv("TELEGRAM_URL", "") + "/bot" + botApiKey
}

func (b Bot) RemoveKeyboard(messageId int, userId string) {
	botApi, err := tgbotapi.NewBotAPI(helper.GetEnv("TELEGRAM_BOT", ""))
	if err != nil {
		return
	}

	user, err := strconv.Atoi(userId)

	if err != nil {
		return
	}

	msg := tgbotapi.NewEditMessageReplyMarkup(int64(user), messageId, tgbotapi.NewInlineKeyboardMarkup())
	msg.ReplyMarkup = nil

	_, err = botApi.Send(msg)

	if err != nil {
		return
	}

	logItem := model.Log{IsBot: true, TelegramId: user, Payload: ""}
	repository.Create(&logItem)
}

func (b Bot) EditMessageKeyboard(messageId int, userId string, keyboard tgbotapi.InlineKeyboardMarkup) {
	botApi, err := tgbotapi.NewBotAPI(helper.GetEnv("TELEGRAM_BOT", ""))
	if err != nil {
		return
	}

	user, err := strconv.Atoi(userId)

	if err != nil {
		return
	}

	msg := tgbotapi.NewEditMessageReplyMarkup(int64(user), messageId, keyboard)

	_, err = botApi.Send(msg)

	if err != nil {
		return
	}

	logItem := model.Log{IsBot: true, TelegramId: user, Payload: ""}
	repository.Create(&logItem)
}

func (b Bot) EditMessage(messageId int, message string, userId string) {
	botApi, err := tgbotapi.NewBotAPI(helper.GetEnv("TELEGRAM_BOT", ""))
	if err != nil {
		return
	}

	user, err := strconv.Atoi(userId)

	if err != nil {
		return
	}

	msg := tgbotapi.NewEditMessageText(int64(user), messageId, message)

	_, err = botApi.Send(msg)

	if err != nil {
		return
	}

	logItem := model.Log{IsBot: true, TelegramId: user, Payload: message}
	repository.Create(&logItem)
}

func (b Bot) SimpleSendMessage(message string, userId string, keyboard interface{}) {
	botApi, err := tgbotapi.NewBotAPI(helper.GetEnv("TELEGRAM_BOT", ""))
	if err != nil {
		return
	}

	user, err := strconv.Atoi(userId)

	if err != nil {
		return
	}

	msg := tgbotapi.NewMessage(int64(user), message)
	msg.ReplyMarkup = keyboard
	msg.ParseMode = "html"

	_, err = botApi.Send(msg)

	if err != nil {
		return
	}

	logItem := model.Log{IsBot: true, TelegramId: user, Payload: message}
	repository.Create(&logItem)
}

func (b Bot) SendPhoto(photoPath, userId string) {
	botApi, err := tgbotapi.NewBotAPI(helper.GetEnv("TELEGRAM_BOT", ""))
	if err != nil {
		return
	}

	photoBytes, err := ioutil.ReadFile(photoPath)
	if err != nil {
		panic(err)
	}
	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photoBytes,
	}

	chatId, err := strconv.Atoi(userId)
	if err != nil {
		return
	}
	_, err = botApi.Send(tgbotapi.NewPhotoUpload(int64(chatId), photoFileBytes))

	payloadData, _ := json.Marshal(mail{ChatId: userId, Text: photoPath, ParseMode: "html"})
	logItem := model.Log{IsBot: true, TelegramId: chatId, Payload: string(payloadData)}
	repository.Create(&logItem)
}

func (b Bot) SendVideo(videoPath, userId string) {
	botApi, err := tgbotapi.NewBotAPI(helper.GetEnv("TELEGRAM_BOT", ""))
	if err != nil {
		return
	}

	videoBytes, err := ioutil.ReadFile(videoPath)
	if err != nil {
		panic(err)
	}
	videoData := tgbotapi.FileBytes{
		Name:  "video",
		Bytes: videoBytes,
	}

	chatId, err := strconv.Atoi(userId)
	if err != nil {
		return
	}
	_, err = botApi.Send(tgbotapi.NewVideoUpload(int64(chatId), videoData))

	payloadData, _ := json.Marshal(mail{ChatId: userId, Text: videoPath, ParseMode: "html"})
	logItem := model.Log{IsBot: true, TelegramId: chatId, Payload: string(payloadData)}
	repository.Create(&logItem)
}

func (b Bot) SendAnimation(animationPath, userId string) {
	botApi, err := tgbotapi.NewBotAPI(helper.GetEnv("TELEGRAM_BOT", ""))
	if err != nil {
		return
	}

	animBytes, err := ioutil.ReadFile(animationPath)
	if err != nil {
		panic(err)
	}
	animationData := tgbotapi.FileBytes{
		Name:  "animation",
		Bytes: animBytes,
	}

	chatId, err := strconv.Atoi(userId)
	if err != nil {
		return
	}
	_, err = botApi.Send(tgbotapi.NewVideoUpload(int64(chatId), animationData))

	payloadData, _ := json.Marshal(mail{ChatId: userId, Text: animationPath, ParseMode: "html"})
	logItem := model.Log{IsBot: true, TelegramId: chatId, Payload: string(payloadData)}
	repository.Create(&logItem)
}

func (b Bot) SetWebhook(url string) []byte {
	if len(url) == 0 {
		url = "https://" + helper.GetEnv("DOMAIN", "") + "/telegram/webhook"
	}

	botApi, err := tgbotapi.NewBotAPI(helper.GetEnv("TELEGRAM_BOT", ""))
	if err != nil {
		log.Println(err)
		return nil
	}

	req, err := botApi.SetWebhook(tgbotapi.NewWebhookWithCert(url, helper.GetEnv("CERT_PATH", "cert.pem")))
	if err != nil {
		log.Println(err)
	}
	return req.Result
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

func (b Bot) SendByMessageType(contentType, messageContent, userId string) {
	content, ok := contentTypeMap[contentType]

	if ok == false {
		return
	}

	if content == "video" {
		b.SendVideo(messageContent, userId)
	} else if content == "photo" {
		b.SendPhoto(messageContent, userId)
	} else if content == "animation" {
		b.SendAnimation(messageContent, userId)
	} else {
		b.SendPhoto(messageContent, userId)
	}
}
