package repository

import (
	"main/config"
	"main/user/model"
)

func CreateJiraUser(user model.JiraUser) {
	config.DbConnection().FirstOrCreate(&user)
}
