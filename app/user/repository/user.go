package repository

import (
	"main/config"
	"main/user/model"
)

func FindOrCreate(user *model.User) {
	nextAction := user.NextAction

	if config.DbConnection().Where("telegram_id = ?", user.TelegramId).Find(&user).RowsAffected == 0 {
		Save(user)
		return
	}

	user.NextAction = nextAction
	Save(user)
}

func Save(user *model.User) {
	config.DbConnection().Save(&user)
}

func Find(user *model.User) {
	config.DbConnection().First(&user)
}

func FindByTelegramId(telegramId int) model.User {
	var user model.User
	config.DbConnection().
		Preload("Permissions.Permission").
		Preload("Settings.Setting").
		Find(&user, &model.User{TelegramId: telegramId})
	return user
}

func GetUsers() []model.User {
	var users []model.User
	config.DbConnection().Model(model.User{}).
		Find(&users)
	return users
}
