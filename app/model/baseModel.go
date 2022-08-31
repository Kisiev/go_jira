package model

import "gorm.io/gorm"

type BaseModel interface {
	create() *gorm.DB
}
