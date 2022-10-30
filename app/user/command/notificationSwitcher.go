package command

import (
	"main/config"
	"main/telegram"
	"main/telegram/entity"
	"main/user/constant"
	"main/user/model"
	"strconv"
	"strings"
)

type NotificationSwitcherCommand struct{}

func (n NotificationSwitcherCommand) Run(update entity.TelegramUpdateInline) {
	var bot telegram.BotInterface = telegram.Bot{}

	user := model.User{TelegramId: update.CallbackQuery.From.Id}
	notificationSetting := model.Setting{Code: constant.NotificationSetting}

	config.DbConnection().First(&user)
	config.DbConnection().First(&notificationSetting)

	userNotificationSetting := model.UserSetting{UserID: int(user.ID), SettingID: int(notificationSetting.ID)}
	config.DbConnection().FirstOrCreate(&userNotificationSetting)

	value := strings.ReplaceAll(update.CallbackQuery.Data, constant.NotificationSetting+"_", "")

	userNotificationSetting.Value = value

	config.DbConnection().Save(&userNotificationSetting)

	bot.RemoveKeyboard(update.CallbackQuery.Message.MessageId, strconv.Itoa(update.CallbackQuery.From.Id))
	bot.SimpleSendMessage("Сохранено", strconv.Itoa(update.CallbackQuery.From.Id), nil)
}

func (n NotificationSwitcherCommand) Support(update entity.TelegramUpdateInline) bool {
	return true
}
