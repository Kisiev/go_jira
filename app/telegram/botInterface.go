package telegram

type BotInterface interface {
	SimpleSendMessage(message string, userId string)
}
