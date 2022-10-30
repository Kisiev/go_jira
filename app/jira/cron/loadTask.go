package cron

import (
	"main/jira"
	"main/jira/repository"
	"main/telegram"
	userModule "main/user/model"
	"strconv"
)

var bot telegram.BotInterface = telegram.Bot{}

func Run() {
	users := repository.JiraUserList()

	for _, user := range users {
		updateTasksForUserAndNotify(user)
	}
}

func updateTasksForUserAndNotify(user userModule.JiraUser) {
	message := jira.LoadAndGetNewTasks(user)

	if len(message) != 0 {
		bot.SimpleSendMessage("Новые задачи\n\n"+message, strconv.Itoa(user.User.TelegramId), nil)
	}
}
