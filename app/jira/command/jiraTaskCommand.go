package command

import (
	"fmt"
	taskFormatter "main/jira/formatter"
	"main/jira/repository"
	"main/telegram"
	telegramEntity "main/telegram/entity"
	"strconv"
)

type JiraTaskCommand struct{}

func (u JiraTaskCommand) Run(update telegramEntity.TelegramUpdate) {
	var bot telegram.BotInterface = telegram.Bot{}

	telegramMessage := update.Message
	user := repository.FindJiraUserByTelegramId(telegramMessage.From.Id)

	tasks := repository.GetUserTask(int64(user.UserID))

	var message string

	for _, item := range tasks {
		message += taskFormatter.FormatMessage(item)
	}

	if len(message) == 0 {
		message = fmt.Sprintf("Для пользователя %s не найдены задачи", user.UserName)
	}

	go bot.SimpleSendMessage(message, strconv.Itoa(telegramMessage.From.Id))
}
