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

func InitDb() *gorm.DB {
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

	err = conn.AutoMigrate(&userModel.User{})
	err = conn.AutoMigrate(&userModel.JiraUser{})
	err = conn.AutoMigrate(&model.Log{})
	err = conn.AutoMigrate(&jiraModel.Task{})

	if err != nil {
		log.Fatal("Cannot migrate")
	}

	return conn
}

func DbConnection() *gorm.DB {
	return InitDb()
}
