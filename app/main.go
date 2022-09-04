package main

import (
	"log"
	"main/telegram/controller"
	"net/http"
	"time"
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
	dateStart := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	dateEnd := time.Now().Format("2006-01-02")
	w.Write([]byte(dateStart + dateEnd))
}
