package main

import (
	"fmt"
	"log"
	"main/config"
	"main/telegram/controller"
	"main/user/model"
	"net/http"
)

func main() {
	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/debug", debug)
	http.HandleFunc("/telegram/setWebhook", controller.SetWebhook)
	http.HandleFunc("/telegram/webhook", controller.Webhook)
	http.HandleFunc("/telegram/getWebhook", controller.GetWebhook)
	err := http.ListenAndServe(":4444", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func debug(w http.ResponseWriter, r *http.Request) {
	var user model.JiraUser
	config.DbConnection().Model(model.JiraUser{}).Joins("join users on users.id = jira_users.user_id").Where("users.telegram_id = ?", "109946632").First(&user)
	F := 1
	fmt.Print(F)
}
