package telegram

import "gopkg.in/telegram-bot-api.v4"

type BotInterface interface {
	SimpleSendMessage(message string, userId string, keyboard interface{})
	RemoveKeyboard(messageId int, message string)
	EditMessageKeyboard(messageId int, userId string, keyboard tgbotapi.InlineKeyboardMarkup)
	EditMessage(messageId int, message string, userId string)
	SetWebhook(url string) []byte
	GetWebhookInfo() []byte
	SendPhoto(photoPath, userId string)
	SendVideo(videoPath, userId string)
	SendAnimation(animationPath, userId string)
	SendByMessageType(contentType, messageContent, userId string)
}
