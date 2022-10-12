package command

import (
	"fmt"
	"github.com/gofrs/uuid"
	"io"
	"main/helper"
	"main/telegram"
	"main/telegram/entity"
	"net/http"
	"os"
	"strconv"
)

type UploadCommand struct{}

func (s UploadCommand) Run(update entity.TelegramUpdate) {
	var bot telegram.BotInterface = telegram.Bot{}
	telegramMessage := update.Message

	fileName, _ := uuid.NewV4()
	fullPath := fmt.Sprintf("%s%s.png", helper.GetEnv("FILES_PATH", ""), fileName.String())

	img, err := os.Open(helper.GetEnv("FILES_PATH", ""))

	if err != nil {
		err = os.MkdirAll(helper.GetEnv("FILES_PATH", ""), 0750)
		if err != nil {
			return
		}
	}

	img, err = os.Create(fullPath)

	if err != nil {
		return
	}

	defer img.Close()

	resp, _ := http.Get(telegramMessage.Text)
	defer resp.Body.Close()

	_, err = io.Copy(img, resp.Body)
	if err != nil {
		return
	}

	bot.SendPhoto(fullPath, strconv.Itoa(update.Message.From.Id))
}
