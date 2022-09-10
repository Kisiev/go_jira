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

func getDb() *gorm.DB {
	connectionString := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		helper.GetEnv("DB_CONNECTION", "postgres"),
		helper.GetEnv("DB_USERNAME", "test"),
		helper.GetEnv("DB_PASSWORD", "test"),
		helper.GetEnv("DB_HOST", "localhost"),
		helper.GetEnv("DB_PORT", "5432"),
		helper.GetEnv("DB_DATABASE", "test"),
	)

	conn, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	return conn
}

func InitDb() {
	conn := getDb()

	conn.AutoMigrate(&userModel.User{})
	conn.AutoMigrate(&userModel.JiraUser{})
	conn.AutoMigrate(&model.Log{})
	conn.AutoMigrate(&jiraModel.Task{})
}

func DbConnection() *gorm.DB {
	return getDb()
}
