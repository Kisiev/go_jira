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

func FindSettingByCode(code string) model.Setting {
	var settingItem model.Setting
	config.DbConnection().Where("code = ?", code).First(&settingItem)
	return settingItem
}
