package model

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	TelegramId int    `gorm:"null;column=telegram_id"`
	IsBot      bool   `gorm:"default=false;column=is_bot"`
	Payload    string `gorm:"null;column=payload"`
}
