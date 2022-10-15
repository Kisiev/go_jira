package cron

import (
	"main/config"
	"main/file"
	"main/jira/repository"
	"main/telegram"
	"main/telegram/model"
	"math/rand"
	"strconv"
	"time"
)

func Motivate() {
	var bot telegram.BotInterface = telegram.Bot{}

	users := repository.JiraUserList()

	for _, user := range users {

		filePath, err := file.GetRandomFilepath()

		if err != nil {
			return
		}

		var motivation []model.Motivation
		config.DbConnection().Model(model.Motivation{}).Where("is_active = ?", true).Find(&motivation)

		if len(motivation) == 0 {
			return
		}

		rand.Seed(time.Now().UnixNano())
		min := 0
		max := len(motivation)
		randomMotivationIndex := rand.Intn(max-min) + min

		bot.SendPhoto(filePath, strconv.Itoa(user.User.TelegramId))
		bot.SimpleSendMessage(motivation[randomMotivationIndex].Title, strconv.Itoa(user.User.TelegramId))
	}
}
