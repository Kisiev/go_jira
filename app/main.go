package main

import (
	"github.com/robfig/cron"
	"log"
	"main/config"
	"main/helper"
	cronCommand "main/jira/cron"
	"main/telegram"
	"main/telegram/controller"
	"net/http"
	"time"
)

func main() {
	config.InitDb()
	setWebhook()
	go cronItems()
	handleRequest()
}

func setWebhook() {
	var bot telegram.BotInterface = telegram.Bot{}
	bot.SetWebhook("")
}

func cronItems() {
	item := cron.New()
	err := item.AddFunc("@every 5m", func() {
		cronCommand.Run()
	})
	if err != nil {
		return
	}
	item.Run()
}

func handleRequest() {
	http.HandleFunc("/debug", debug)
	http.HandleFunc("/telegram/setWebhook", controller.SetWebhook)
	http.HandleFunc("/telegram/webhook", controller.Webhook)
	http.HandleFunc("/telegram/getWebhook", controller.GetWebhook)
	//err := http.ListenAndServe(":8081", nil)
	err := http.ListenAndServeTLS(":443",
		helper.GetEnv("CERT_PATH", "cert.pem"),
		helper.GetEnv("KEY_PATH", "key.pem"),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func debug(w http.ResponseWriter, r *http.Request) {
	dateStart := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	dateEnd := time.Now().Format("2006-01-02")
	w.Write([]byte(dateStart + dateEnd))
}
