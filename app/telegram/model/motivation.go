package model

import "gorm.io/gorm"

type Motivation struct {
	gorm.Model
	Title    string `gorm:"null;column=title"`
	isActive bool   `gorm:"default=true;column=is_active"`
}
