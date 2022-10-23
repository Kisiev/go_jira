package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	ID       int64  `gorm:"primaryKey;autoIncrement:true"`
	Path     string `gorm:"not null;column:path"`
	Type     string `gorm:"not null;column:type"`
	IsActive bool   `gorm:"not null;default:true;column:is_active"`
}
