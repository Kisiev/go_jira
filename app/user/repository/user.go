package repository

import (
	"main/config"
	"main/user/model"
)

func FindOrCreate(user *model.User) {
	config.DbConnection().FirstOrCreate(&user)
}

func Find(user *model.User) {
	config.DbConnection().First(&user)
}

func FindByTelegramId(telegramId int) model.User {
	var user model.User
	config.DbConnection().Find(&user, &model.User{TelegramId: telegramId})
	return user
}

func FindJiraUserByTelegramId(telegramId int) model.JiraUser {
	var user model.JiraUser
	config.DbConnection().Model(model.JiraUser{}).Joins("join users on users.id = jira_users.user_id").Where("users.telegram_id = ?", telegramId).First(&user)
	return user
}
