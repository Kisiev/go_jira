package command

import (
	"fmt"
	"github.com/gofrs/uuid"
	"io"
	"main/helper"
	"main/telegram"
	"main/telegram/entity"
	"main/user"
	"main/user/model"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type UploadCommand struct{}

func (u UploadCommand) Run(update entity.TelegramUpdate) {
	var bot telegram.BotInterface = telegram.Bot{}
	telegramMessage := update.Message

	urls := strings.Split(telegramMessage.Text, "\n")
	var savedPictures int

	for _, url := range urls {
		url = strings.TrimSpace(url)

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
			continue
		}

		resp, err := http.Get(url)

		if err != nil {
			continue
		}

		_, err = io.Copy(img, resp.Body)
		if err != nil {
			continue
		}

		img.Close()
		resp.Body.Close()

		savedPictures++
	}

	bot.SimpleSendMessage(fmt.Sprintf("Сохранено картинок %d", savedPictures), strconv.Itoa(update.Message.From.Id))
}

func (u UploadCommand) Support(update entity.TelegramUpdate) bool {
	var bot telegram.BotInterface = telegram.Bot{}

	telegramUser := strconv.Itoa(update.Message.From.Id)
	if user.CheckPermission(telegramUser, model.PermissionCanUpload) {
		return true
	}

	bot.SimpleSendMessage("Нет доступа", telegramUser)
	return false
}
