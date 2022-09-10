package cron

import (
	"fmt"
	"main/jira"
	taskFormatter "main/jira/formatter"
	"main/jira/model"
	"main/jira/repository"
	"main/telegram"
	userModule "main/user/model"
	"strconv"
)

var bot telegram.BotInterface = telegram.Bot{}

func Run() {
	users := repository.JiraUserList()

	for _, user := range users {
		go updateTasksForUserAndNotify(user)
	}
}

func updateTasksForUserAndNotify(user userModule.JiraUser) {
	rawJiraFilter := fmt.Sprintf("project = TRACEWAY and assignee=%s and status not in (Закрыто, Выполнено, Done, CLOSED, Canceled) ORDER BY priority DESC, created DESC", user.UserName)
	jiraData := jira.GetTasksForUser(rawJiraFilter)

	var formatter taskFormatter.Formatter = taskFormatter.JiraFormatter{}
	data := formatter.Format(jiraData)

	var message string
	var notDeletingTaskUrls []string
	for _, item := range data {
		newTask := model.Task{Title: item.Title, Url: item.Url, Priority: item.Priority, Status: item.Status, UserId: user.UserID}
		if repository.CheckIfExist(newTask) == 0 {
			message += taskFormatter.FormatMessage(newTask)
		}
		repository.CreateIfNotExistTask(&newTask)
		notDeletingTaskUrls = append(notDeletingTaskUrls, newTask.Url)
	}

	repository.DeleteTasksWithout(user.UserID, notDeletingTaskUrls)

	if len(message) != 0 {
		bot.SimpleSendMessage("Новые задачи\n\n"+message, strconv.Itoa(user.User.TelegramId))
	}
}
