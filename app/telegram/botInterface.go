package telegram

type BotInterface interface {
	SimpleSendMessage(message string, userId string)
	SetWebhook(url string) []byte
	GetWebhookInfo() []byte
	SendPhoto(photoPath, userId string)
	SendVideo(videoPath, userId string)
	SendAnimation(animationPath, userId string)
	SendByMessageType(contentType, messageContent, userId string)
}
