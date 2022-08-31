package main

import (
	"log"
	"main/config"
	taskFormatter "main/formatter"
	"main/helper"
	"main/model"
	"main/telegram"
	"net/http"
)

var bot telegram.BotInterface

func main() {
	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/task", taskList)
	http.HandleFunc("/new-task", newTaskList)
	err := http.ListenAndServe(":4444", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func newTaskList(w http.ResponseWriter, r *http.Request) {
	jiraData := getData()

	var formatter taskFormatter.Formatter = taskFormatter.JiraFormatter{}
	data := formatter.Format(jiraData)

	bot = telegram.Bot{}
	var message string

	for _, item := range data {
		taskItem := model.Task{}
		config.DbConnection().Find(&taskItem, model.Task{Url: item.Url})

		if len(taskItem.Title) == 0 {
			message += taskFormatter.FormatMessage(item)
		}
	}

	if len(message) > 0 {
		go bot.SimpleSendMessage("Новые задачи\n\n"+message, helper.GetEnv("TELEGRAM_ID", ""))
	}
}

func taskList(w http.ResponseWriter, r *http.Request) {
	jiraData := getData()

	var formatter taskFormatter.Formatter = taskFormatter.JiraFormatter{}
	data := formatter.Format(jiraData)

	bot = telegram.Bot{}
	var message string

	var notFoundTasks []model.Task
	var allTask []model.Task

	for _, item := range data {
		taskItem := model.Task{}
		config.DbConnection().Find(&taskItem, model.Task{Url: item.Url})

		if len(taskItem.Title) == 0 {
			config.DbConnection().Create(&item)
			notFoundTasks = append(notFoundTasks, item)
		}

		allTask = append(allTask, item)
		message += taskFormatter.FormatMessage(item)
	}

	var urls []string

	for _, item := range allTask {
		urls = append(urls, item.Url)
	}

	config.DbConnection().Where("url NOT IN (?)", urls).Delete(model.Task{})

	go bot.SimpleSendMessage(message, helper.GetEnv("TELEGRAM_ID", ""))
}
