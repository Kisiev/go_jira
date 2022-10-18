package config

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"main/helper"
	jiraModel "main/jira/model"
	"main/telegram/model"
	userModel "main/user/model"
)

var DB *gorm.DB

func getDb() {
	connectionString := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		helper.GetEnv("DB_CONNECTION", "postgres"),
		helper.GetEnv("DB_USERNAME", "test"),
		helper.GetEnv("DB_PASSWORD", "test"),
		helper.GetEnv("DB_HOST", "localhost"),
		helper.GetEnv("DB_PORT", "5432"),
		helper.GetEnv("DB_DATABASE", "test"),
	)

	var err error

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
}

func InitDb() {
	getDb()

	DB.AutoMigrate(&userModel.User{})
	DB.AutoMigrate(&userModel.JiraUser{})
	DB.AutoMigrate(&model.Log{})
	DB.AutoMigrate(&jiraModel.Task{})
	DB.AutoMigrate(&model.Motivation{})
}

func DbConnection() *gorm.DB {
	return DB
}
