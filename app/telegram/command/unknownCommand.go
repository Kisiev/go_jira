package command

import (
	"fmt"
	"main/helper"
	"main/telegram"
	"main/telegram/entity"
)

type UnknownCommand struct{}

func (u UnknownCommand) Run(update entity.TelegramUpdate) {
	var bot telegram.BotInterface = telegram.Bot{}

	adminId := helper.GetEnv("TELEGRAM_ID", "")

	message := fmt.Sprintf("Пользователь: %d\nИмя: %s\nСообщение: %s\n", update.Message.From.Id, update.Message.From.LastName+" "+update.Message.From.FirstName, update.Message.Text)
	bot.SimpleSendMessage(message, adminId, nil)
}

func (u UnknownCommand) Support(update entity.TelegramUpdate) bool {
	return true
}
