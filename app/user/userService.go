package user

import (
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
