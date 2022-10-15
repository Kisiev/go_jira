package command

import (
	"main/file"
	"main/telegram"
	"main/telegram/entity"
	"strconv"
)

type RandomFileCommand struct{}

func (s RandomFileCommand) Run(update entity.TelegramUpdate) {
	var bot telegram.BotInterface = telegram.Bot{}

	fullPath, err := file.GetRandomFilepath()

	if err != nil {
		return
	}

	bot.SendPhoto(fullPath, strconv.Itoa(update.Message.From.Id))
}
