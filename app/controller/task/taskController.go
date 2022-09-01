package task

import (
	taskFormatter "main/formatter"
	"main/helper"
	"main/jira"
	"main/telegram"
	"net/http"
)

var bot telegram.BotInterface

func List(w http.ResponseWriter, r *http.Request) {
	jiraData := jira.GetTasks()

	var formatter taskFormatter.Formatter = taskFormatter.JiraFormatter{}
	data := formatter.Format(jiraData)

	bot = telegram.Bot{}
	var message string

	for _, item := range data {
		message += taskFormatter.FormatMessage(item)
	}

	go bot.SimpleSendMessage(message, helper.GetEnv("TELEGRAM_ID", ""))
}
