package jira

import (
	"fmt"
	taskFormatter "main/jira/formatter"
	"main/jira/model"
	"main/jira/repository"
	userModule "main/user/model"
)

func LoadAndGetNewTasks(user userModule.JiraUser) string {
	rawJiraFilter := fmt.Sprintf("assignee=%s and status not in (Закрыто, Выполнено, Done, CLOSED, Canceled) ORDER BY priority DESC, created DESC", user.UserName)
	jiraData := GetTasksForUser(rawJiraFilter)

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

	return message
}
