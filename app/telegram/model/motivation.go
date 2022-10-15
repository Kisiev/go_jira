package model

import "gorm.io/gorm"

type Motivation struct {
	gorm.Model
	Title    string `gorm:"null;column:title"`
	IsActive bool   `gorm:"not null;default:true;column:is_active"`
}
