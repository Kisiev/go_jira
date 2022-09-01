package model

type JiraUser struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserName string `gorm:"not null"`
	UserID   int    `gorm:"column:user_id"`
	User     *User  `gorm:"foreignKey:ID;references:UserID"`
}
