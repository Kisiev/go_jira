package command

import (
	fileRepository "main/file/repository"
	"main/telegram"
	"main/telegram/entity"
	"main/user"
	"main/user/model"
	userRepository "main/user/repository"
	"strconv"
)

type RandomFileCommand struct{}

func (r RandomFileCommand) Run(update entity.TelegramUpdate) {
	var bot telegram.BotInterface = telegram.Bot{}

	userModel := userRepository.FindByTelegramId(update.Message.From.Id)
	fileCount := fileRepository.GetRandPathForUser(int(userModel.ID))
	fileModel := fileRepository.GetById(fileCount.ID)
	fileRepository.AddCountToFileForUser(int(fileModel.ID), int(userModel.ID), fileCount.Count+1)

	bot.SendByMessageType(fileModel.Type, fileModel.Path, strconv.Itoa(update.Message.From.Id))
}

func (r RandomFileCommand) Support(update entity.TelegramUpdate) bool {
	var bot telegram.BotInterface = telegram.Bot{}

	telegramUser := strconv.Itoa(update.Message.From.Id)
	if user.CheckPermission(telegramUser, model.PermissionRandMode) {
		return true
	}

	bot.SimpleSendMessage("Нет доступа", telegramUser, nil)
	return false
}
