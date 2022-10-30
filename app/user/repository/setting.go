package repository

import (
	"main/config"
	"main/user/model"
)

func AllSettings() []model.Setting {
	var settings []model.Setting
	config.DbConnection().Find(&settings)
	return settings
}
