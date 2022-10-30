package command

import (
	"gopkg.in/telegram-bot-api.v4"
	"main/telegram"
	"main/telegram/entity"
	"main/user/constant"
	"strconv"
)

type SettingCommand struct{}

func (s SettingCommand) Run(update entity.TelegramUpdateInline) {
	var bot telegram.BotInterface = telegram.Bot{}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Вкл", constant.NotificationSetting+"_1"),
			tgbotapi.NewInlineKeyboardButtonData("Выкл", constant.NotificationSetting+"_0"),
		),
	)

	bot.EditMessageKeyboard(update.CallbackQuery.Message.MessageId, strconv.Itoa(update.CallbackQuery.From.Id), keyboard)
}

func (s SettingCommand) Support(update entity.TelegramUpdateInline) bool {
	return true
}
