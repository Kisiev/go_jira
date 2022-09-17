package command

import (
	"encoding/json"
	"main/jira"
	jiraRepository "main/jira/repository"
	"main/jira/validator"
	"main/telegram"
	"main/telegram/entity"
	"main/user/model"
	userRepository "main/user/repository"
	"strconv"
)

type UserCreateCommand struct{}

func (u UserCreateCommand) Run(update entity.TelegramUpdate) {
	message := update.Message

	jiraUserName := message.Text

	var bot telegram.BotInterface = telegram.Bot{}

	user := userRepository.FindByTelegramId(message.From.Id)
	jiraUser := model.JiraUser{UserName: jiraUserName, UserID: int(user.ID)}
	err := validator.Validate(jiraUser)

	if err != nil {
		go bot.SimpleSendMessage(err.Error(), strconv.Itoa(user.TelegramId))
		return
	}

	jiraRepository.CreateJiraUser(jiraUser)

	nextActionStr, err := json.Marshal(model.NextAction{})
	user.NextAction = string(nextActionStr)
	userRepository.Save(&user)

	jira.LoadAndGetNewTasks(jiraUser)

	go bot.SimpleSendMessage("Пользователь добавлен", strconv.Itoa(user.TelegramId))
}
