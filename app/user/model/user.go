package model

type User struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement:true;column=id"`
	Name       string `gorm:"not null;column=name"`
	TelegramId int    `gorm:"null;column=telegram_id"`
	NextAction string `gorm:"null;column=next_action"`
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
