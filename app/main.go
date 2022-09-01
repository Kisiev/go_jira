package main

import (
	"log"
	"main/config"
	"main/controller/task"
	"main/controller/telegram"
	"net/http"
)

func main() {
	config.InitDb()
	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/task", task.List)
	http.HandleFunc("/telegram/setWebhook", telegram.SetWebhook)
	http.HandleFunc("/telegram/webhook", telegram.Webhook)
	http.HandleFunc("/telegram/getWebhook", telegram.GetWebhook)
	err := http.ListenAndServe(":4444", nil)
	if err != nil {
		log.Fatal(err)
	}
}
