package command

import (
	"fmt"
	"github.com/gofrs/uuid"
	"io"
	"io/ioutil"
	fileModel "main/file/model"
	fileRepository "main/file/repository"
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

var allowedExtension = map[string]string{
	"image/jpeg": "jpeg",
	"image/png":  "png",
	"image/gif":  "gif",
	"video/mp4":  "mp4",
}

func (u UploadCommand) Run(update entity.TelegramUpdate) {
	var bot telegram.BotInterface = telegram.Bot{}
	telegramMessage := update.Message

	urls := strings.Split(telegramMessage.Text, "\n")
	var savedPictures int

	for _, url := range urls {
		url = strings.TrimSpace(url)

		resp, err := http.Get(url)

		if err != nil {
			continue
		}

		imageContent, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		mimeType := http.DetectContentType(imageContent)
		extensionTxt, err := getAllowedExtensionByMimeType(mimeType)

		if err != nil {
			bot.SimpleSendMessage(err.Error()+" "+mimeType, strconv.Itoa(update.Message.From.Id), nil)
			continue
		}

		fileName, _ := uuid.NewV4()
		fullPath := fmt.Sprintf("%s%s.%s", helper.GetEnv("FILES_PATH", ""), fileName.String(), extensionTxt)

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

		fileSize, err := io.WriteString(img, string(imageContent))

		if err != nil || fileSize == 0 {
			continue
		}

		img.Close()
		resp.Body.Close()

		fileItem := fileModel.File{Path: fullPath, Type: mimeType, IsActive: true}
		fileRepository.Create(&fileItem)

		savedPictures++
	}

	bot.SimpleSendMessage(fmt.Sprintf("Сохранено картинок %d", savedPictures), strconv.Itoa(update.Message.From.Id), nil)
}

func getAllowedExtensionByMimeType(mimetype string) (string, error) {
	if value, ok := allowedExtension[mimetype]; ok {
		return value, nil
	}

	return "", fmt.Errorf("недопустимый формат")
}

func (u UploadCommand) Support(update entity.TelegramUpdate) bool {
	var bot telegram.BotInterface = telegram.Bot{}

	telegramUser := strconv.Itoa(update.Message.From.Id)
	if user.CheckPermission(telegramUser, model.PermissionCanUpload) {
		return true
	}

	bot.SimpleSendMessage("Нет доступа", telegramUser, nil)
	return false
}
