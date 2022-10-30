package model

import "gorm.io/gorm"

type Setting struct {
	gorm.Model
	ID    uint64 `gorm:"primaryKey;autoIncrement:true;column:id"`
	Title string `gorm:"not null;column:title"`
	Code  string `gorm:"not null;column:code"`
}

type UserSetting struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey;autoIncrement:true;column:id"`
	UserID    int    `gorm:"column:user_id"`
	User      User
	SettingID int `gorm:"column:setting_id"`
	Setting   Setting
	Value     string `gorm:"column:value"`
}
