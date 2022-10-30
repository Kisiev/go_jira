package user

import (
	"main/user/constant"
	"main/user/repository"
	"strconv"
)

func CheckPermission(telegramUserId string, permission string) bool {
	telegramUser, _ := strconv.Atoi(telegramUserId)
	user := repository.FindByTelegramId(telegramUser)

	for _, item := range user.Permissions {
		if item.Permission.Name == permission {
			return true
		}
	}

	return false
}

func CanNotification(telegramUserId string) bool {
	telegramUser, _ := strconv.Atoi(telegramUserId)
	user := repository.FindByTelegramId(telegramUser)

	for _, item := range user.Settings {
		if item.Setting.Code == constant.NotificationSetting && item.Value == "0" {
			return false
		}
	}

	return true
}
