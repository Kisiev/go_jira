package command

import (
	"main/config"
	"main/telegram"
	"main/telegram/entity"
	"main/user/constant"
	"main/user/model"
	"main/user/repository"
	"strconv"
	"strings"
)

type NotificationSwitcherCommand struct{}

func (n NotificationSwitcherCommand) Run(update entity.TelegramUpdateInline) {
	var bot telegram.BotInterface = telegram.Bot{}

	user := repository.FindByTelegramId(update.CallbackQuery.From.Id)
	notificationSetting := repository.FindSettingByCode(constant.NotificationSetting)

	userNotificationSetting := model.UserSetting{UserID: int(user.ID), SettingID: int(notificationSetting.ID)}
	if config.DbConnection().Where("user_id = ? AND setting_id = ?", user.ID, notificationSetting.ID).
		First(&userNotificationSetting).RowsAffected == 0 {
		config.DbConnection().Save(&userNotificationSetting)
	}

	value := strings.ReplaceAll(update.CallbackQuery.Data, constant.NotificationSetting+"_", "")

	userNotificationSetting.Value = value

	config.DbConnection().Save(&userNotificationSetting)

	bot.RemoveKeyboard(update.CallbackQuery.Message.MessageId, strconv.Itoa(update.CallbackQuery.From.Id))
	bot.SimpleSendMessage("Сохранено", strconv.Itoa(update.CallbackQuery.From.Id), nil)
}

func (n NotificationSwitcherCommand) Support(update entity.TelegramUpdateInline) bool {
	return true
}
