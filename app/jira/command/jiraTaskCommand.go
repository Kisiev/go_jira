package command

import (
	"fmt"
	"main/jira"
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

	rawJiraFilter := fmt.Sprintf("project = TRACEWAY and assignee=%s and status not in (Закрыто, Выполнено, Done, CLOSED, Canceled) ORDER BY created DESC", user.UserName)
	jiraData := jira.GetTasksForUser(rawJiraFilter)

	var formatter taskFormatter.Formatter = taskFormatter.JiraFormatter{}
	data := formatter.Format(jiraData)

	var message string

	for _, item := range data {
		message += taskFormatter.FormatMessage(item)
	}

	if len(message) == 0 {
		message = fmt.Sprintf("Для пользователя %s не найдены задачи", user.UserName)
	}

	go bot.SimpleSendMessage(message, strconv.Itoa(telegramMessage.From.Id))
}
