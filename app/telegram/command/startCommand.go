package command

import (
	"encoding/json"
	"main/telegram"
	"main/telegram/constant"
	"main/telegram/entity"
	"main/user/model"
	"main/user/repository"
	"strconv"
)

type StartCommand struct{}

func (s StartCommand) Run(update entity.TelegramUpdate) {
	var bot telegram.BotInterface = telegram.Bot{}
	message := update.Message

	existUser := repository.FindByTelegramId(message.From.Id)

	if existUser.ID > 0 {
		bot.SimpleSendMessage("Пользователь уже зарегистрирован", strconv.Itoa(message.From.Id), nil)
		return
	}

	fullName := message.From.FirstName + " " + message.From.LastName
	nextAction := model.NextAction{Action: constant.ActionSetJiraUserName}
	nextActionStr, err := json.Marshal(nextAction)

	if err != nil {
		bot.SimpleSendMessage("Ошибка обработки"+err.Error(), strconv.Itoa(message.From.Id), nil)
	}

	user := model.User{Name: fullName, TelegramId: message.From.Id, NextAction: string(nextActionStr)}
	repository.FindOrCreate(&user)

	bot.SimpleSendMessage("Пользователь зарегистрирован. Введите имя пользователя Jira", strconv.Itoa(user.TelegramId), nil)
}

func (s StartCommand) Support(update entity.TelegramUpdate) bool {
	return true
}
