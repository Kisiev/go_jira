package cron

import (
	"main/config"
	fileRepository "main/file/repository"
	"main/telegram"
	"main/telegram/model"
	userModule "main/user"
	userModel "main/user/model"
	"main/user/repository"
	"math/rand"
	"strconv"
	"time"
)

func Motivate() {
	var bot telegram.BotInterface = telegram.Bot{}

	users := repository.GetUsers()

	var motivation []model.Motivation
	config.DbConnection().Model(model.Motivation{}).Where("is_active = ?", true).Find(&motivation)

	if len(motivation) == 0 {
		return
	}

	for _, user := range users {
		if !userModule.CheckPermission(strconv.Itoa(user.TelegramId), userModel.PermissionFunNotification) {
			continue
		}

		if !userModule.CanNotification(strconv.Itoa(user.TelegramId)) {
			continue
		}

		fileCount := fileRepository.GetRandPathForUser(int(user.ID))
		fileModel := fileRepository.GetById(fileCount.ID)
		fileRepository.AddCountToFileForUser(int(fileModel.ID), int(user.ID), fileCount.Count+1)

		rand.Seed(time.Now().UnixNano())
		min := 0
		max := len(motivation)
		randomMotivationIndex := rand.Intn(max-min) + min

		bot.SendByMessageType(fileModel.Type, fileModel.Path, strconv.Itoa(user.TelegramId))
		bot.SimpleSendMessage(motivation[randomMotivationIndex].Title, strconv.Itoa(user.TelegramId), nil)
	}
}
