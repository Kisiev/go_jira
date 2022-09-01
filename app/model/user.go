package model

type User struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name       string `gorm:"not null"`
	TelegramId int    `gorm:"null"`
}
