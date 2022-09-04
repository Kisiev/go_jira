package model

type Log struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement:true;column=id"`
	TelegramId int    `gorm:"null;column=telegram_id"`
	NextAction string `gorm:"null;column=next_action"`
}
