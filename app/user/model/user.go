package model

import "gorm.io/gorm"

const PermissionFunNotification string = "fun-notification"
const PermissionRandMode string = "rand-mode"
const PermissionCanUpload string = "can-upload"

type User struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement:true;column:id"`
	Name        string `gorm:"not null;column:name"`
	TelegramId  int    `gorm:"null;column:telegram_id"`
	NextAction  string `gorm:"null;column:next_action"`
	Permissions []UsersPermission
	Settings    []UserSetting
}

type JiraUser struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserName string `gorm:"not null"`
	UserID   int    `gorm:"column:user_id"`
	User     User
}

type NextAction struct {
	Action  string `json:"action"`
	Payload struct {
		RawData      string `json:"raw_data"`
		ActionParams string `json:"action_params"`
	} `json:"payload"`
}

type Permission struct {
	gorm.Model
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
}

type UsersPermission struct {
	gorm.Model
	UserID       int `gorm:"column:user_id"`
	User         User
	PermissionID int `gorm:"column:permission_id"`
	Permission   Permission
}
