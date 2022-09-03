package repository

import (
	"main/config"
	"main/user/model"
)

func FindOrCreate(user *model.User) {
	config.DbConnection().FirstOrCreate(&user)
}

func Save(user *model.User) {
	config.DbConnection().Save(&user)
}

func Find(user *model.User) {
	config.DbConnection().First(&user)
}

func FindByTelegramId(telegramId int) model.User {
	var user model.User
	config.DbConnection().Find(&user, &model.User{TelegramId: telegramId})
	return user
}
