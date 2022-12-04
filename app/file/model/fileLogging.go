package model

import (
	"gorm.io/gorm"
	"main/user/model"
)

type FileLogging struct {
	gorm.Model
	ID     int64 `gorm:"primaryKey;autoIncrement:true"`
	UserID int   `gorm:"not null;column:user_id"`
	FileID int   `gorm:"not null;column:file_id"`
	Count  int   `gorm:"not null;column:count;default:0"`
	User   model.User
	File   File
}
