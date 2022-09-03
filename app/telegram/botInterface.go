package telegram

type BotInterface interface {
	SimpleSendMessage(message string, userId string)
	SetWebhook(url string) []byte
	GetWebhookInfo() []byte
}
