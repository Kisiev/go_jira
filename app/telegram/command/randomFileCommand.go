package command

import (
	"main/config"
	"main/file"
	fileModel "main/file/model"
	"main/telegram"
	"main/telegram/entity"
	"main/user"
	"main/user/model"
	"strconv"
)

type RandomFileCommand struct{}

func (r RandomFileCommand) Run(update entity.TelegramUpdate) {
	var bot telegram.BotInterface = telegram.Bot{}

	var files []fileModel.File
	config.DbConnection().Find(&files)

	fileItem, err := file.GetRandomFilepath(files)

	if err != nil {
		return
	}

	bot.SendByMessageType(fileItem.Type, fileItem.Path, strconv.Itoa(update.Message.From.Id))
}

func (r RandomFileCommand) Support(update entity.TelegramUpdate) bool {
	var bot telegram.BotInterface = telegram.Bot{}

	telegramUser := strconv.Itoa(update.Message.From.Id)
	if user.CheckPermission(telegramUser, model.PermissionRandMode) {
		return true
	}

	bot.SimpleSendMessage("Нет доступа", telegramUser)
	return false
}
