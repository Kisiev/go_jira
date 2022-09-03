package repository

import (
	"main/config"
	"main/user/model"
)

func CreateJiraUser(user model.JiraUser) {
	if config.DbConnection().Model(&user).Where("user_id = ?", user.UserID).Update("user_name", user.UserName).RowsAffected == 0 {
		config.DbConnection().Create(&user)
	}
}

func FindJiraUserByTelegramId(telegramId int) model.JiraUser {
	var user model.JiraUser
	config.DbConnection().Model(model.JiraUser{}).Joins("join users on users.id = jira_users.user_id").Where("users.telegram_id = ?", telegramId).First(&user)
	return user
}
