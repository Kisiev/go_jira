package main

import (
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"main/config"
	"main/file/model"
	"main/helper"
	cronCommand "main/jira/cron"
	"main/telegram"
	"main/telegram/controller"
	telegramCron "main/telegram/cron"
	"net/http"
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

	err = item.AddFunc("0 0 8-14 * * 1-5", func() {
		telegramCron.Motivate()
	})

	if err != nil {
		return
	}

	item.Start()
}

func handleRequest() {
	http.HandleFunc("/debugg", debugg)
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

func debugg(w http.ResponseWriter, r *http.Request) {
	files, _ := ioutil.ReadDir(helper.GetEnv("FILES_PATH", ""))

	for _, file := range files {
		fileContent, err := ioutil.ReadFile(helper.GetEnv("FILES_PATH", "") + file.Name())
		if err != nil {
			return
		}

		mimeType := http.DetectContentType(fileContent)

		fileItem := model.File{Path: helper.GetEnv("FILES_PATH", "") + file.Name(), Type: mimeType, IsActive: true}
		config.DbConnection().Create(&fileItem)
	}
}
