package command

import (
	"io/ioutil"
	"main/helper"
	"main/telegram"
	"main/telegram/entity"
	"math/rand"
	"strconv"
	"time"
)

type RandomFileCommand struct{}

func (s RandomFileCommand) Run(update entity.TelegramUpdate) {
	var bot telegram.BotInterface = telegram.Bot{}

	files, err := ioutil.ReadDir(helper.GetEnv("FILES_PATH", ""))
	if err != nil {
		return
	}

	fileCount := len(files)

	if fileCount == 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := fileCount
	randomFile := rand.Intn(max-min) + min

	fullPath := helper.GetEnv("FILES_PATH", "") + files[randomFile].Name()

	bot.SendPhoto(fullPath, strconv.Itoa(update.Message.From.Id))
}
