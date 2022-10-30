package command

import (
	"gopkg.in/telegram-bot-api.v4"
	"main/telegram"
	"main/telegram/entity"
	"main/user/repository"
	"strconv"
)

type SettingCommand struct{}

func (s SettingCommand) Run(update entity.TelegramUpdate) {
	var bot telegram.BotInterface = telegram.Bot{}

	settings := repository.AllSettings()

	var keyboard tgbotapi.InlineKeyboardMarkup
	var buttons []tgbotapi.InlineKeyboardButton

	for _, item := range settings {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(item.Title, item.Code))
	}

	if len(buttons) <= 0 {
		return
	}

	keyboard = tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(buttons...))
	bot.SimpleSendMessage("Настройки", strconv.Itoa(update.Message.From.Id), keyboard)
}

func (s SettingCommand) Support(update entity.TelegramUpdate) bool {
	return true
}
