package command

import (
	jiraRepository "main/jira/repository"
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
	jiraRepository.CreateJiraUser(jiraUser)

	go bot.SimpleSendMessage("Пользователь добавлен", strconv.Itoa(user.TelegramId))
}
