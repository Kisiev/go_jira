package config

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"main/helper"
	"main/model"
)

var connection *gorm.DB

func InitDb() {
	connectionString := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		helper.GetEnv("DB_CONNECTION", "postgres"),
		helper.GetEnv("DB_USERNAME", "test"),
		helper.GetEnv("DB_PASSWORD", "test"),
		helper.GetEnv("DB_HOST", "localhost"),
		helper.GetEnv("DB_PORT", "5432"),
		helper.GetEnv("DB_DATABASE", "test"),
	)

	connection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	err = connection.AutoMigrate(&model.User{})
	err = connection.AutoMigrate(&model.JiraUser{})

	if err != nil {
		log.Fatal("Cannot migrate")
	}
}

func DbConnection() *gorm.DB {
	return connection
}
