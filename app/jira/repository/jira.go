package repository

import (
	"main/config"
	jiraModel "main/jira/model"
	"main/user/model"
)

func CreateJiraUser(user model.JiraUser) {
	if config.DbConnection().Model(&user).Where("user_id = ?", user.UserID).Update("user_name", user.UserName).RowsAffected == 0 {
		config.DbConnection().Create(&user)
	}
}

func FindJiraUserByTelegramId(telegramId int) model.JiraUser {
	var user model.JiraUser
	config.DbConnection().Model(model.JiraUser{}).Preload("User").Joins("join users on users.id = jira_users.user_id").Where("users.telegram_id = ?", telegramId).First(&user)
	return user
}

func JiraUserList() []model.JiraUser {
	var users []model.JiraUser
	config.DbConnection().Model(model.JiraUser{}).
		Joins("join users on users.id = jira_users.user_id").
		Preload("User").
		Find(&users)
	return users
}

func CheckIfExist(task jiraModel.Task) int64 {
	return config.DbConnection().Where("user_id = ? AND url = ?", task.UserId, task.Url).First(&task).RowsAffected
}

func CreateIfNotExistTask(task *jiraModel.Task) {
	if config.DbConnection().Model(jiraModel.Task{}).
		Where("user_id = ? AND url = ?", task.UserId, task.Url).
		Updates(&task).RowsAffected == 0 {
		config.DbConnection().Create(task)
	}
}

func GetUserTask(userId int64) []jiraModel.Task {
	var tasks []jiraModel.Task
	config.DbConnection().Order("priority desc").Where("user_id = ?", userId).Find(&tasks)
	return tasks
}

func DeleteTasksWithout(userId int, urls []string) {
	config.DbConnection().Where("url not in (?) and user_id = ?", urls, userId).Delete(jiraModel.Task{})
}
