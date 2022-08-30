package main

import (
	taskFormatter "main/formatter"
	"main/helper"
	"main/telegram"
	"net/http"
)

var bot telegram.BotInterface

func main() {
	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	jiraData := getData()

	var formatter taskFormatter.Formatter = taskFormatter.JiraFormatter{}
	data := formatter.Format(jiraData)

	bot = telegram.Bot{}
	var message string

	for _, item := range data {
		message += taskFormatter.FormatMessage(item)
	}

	go bot.SimpleSendMessage(message, helper.GetEnv("TELEGRAM_ID", ""))
}
