package model

type Task struct {
	Id       int64  `gorm:"primaryKey;autoIncrement"`
	Title    string `gorm:"not null"`
	Url      string `gorm:"not null"`
	Priority string `gorm:"not null"`
	Status   string `gorm:"not null"`
}
