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

var bot telegram.BotInterface = telegram.Bot{}

func (s StartCommand) Run(update entity.TelegramUpdate) {
	message := update.Message

	fullName := message.From.FirstName + " " + message.From.LastName
	nextAction := model.NextAction{Action: constant.ActionSetJiraUserName}
	nextActionStr, err := json.Marshal(nextAction)

	if err != nil {
		bot.SimpleSendMessage("Ошибка обработки"+err.Error(), strconv.Itoa(message.From.Id))
	}

	user := model.User{Name: fullName, TelegramId: message.From.Id, NextAction: string(nextActionStr)}
	repository.FindOrCreate(&user)

	bot.SimpleSendMessage("Пользователь зарегистрирован. Введите имя пользователя Jira", strconv.Itoa(user.TelegramId))
}
