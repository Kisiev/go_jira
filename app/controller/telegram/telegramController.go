package telegram

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"main/helper"
	"net/http"
)

func Webhook(w http.ResponseWriter, r *http.Request) {
	all, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	marshal, err := json.Marshal(all)
	if err != nil {
		return
	}
	_, err = base64.StdEncoding.DecodeString(string(marshal))
	if err != nil {
		return
	}
	_ = []byte("sa")
	fmt.Println(marshal)
}

func SetWebhook(w http.ResponseWriter, r *http.Request) {
	webhookUrl := r.URL.Query().Get("url")

	url := fmt.Sprintf("%s/bot%s/setWebhook?url=%s",
		helper.GetEnv("TELEGRAM_URL", ""),
		helper.GetEnv("TELEGRAM_BOT", ""),
		webhookUrl,
	)

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	w.Write(body)
}

func GetWebhook(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/bot%s/getWebhookInfo",
		helper.GetEnv("TELEGRAM_URL", ""),
		helper.GetEnv("TELEGRAM_BOT", ""),
	)

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	w.Write(body)
}
