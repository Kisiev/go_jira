package repository

import (
	"main/config"
	"main/telegram/model"
)

func Create(log *model.Log) {
	config.DbConnection().Save(&log)
}
