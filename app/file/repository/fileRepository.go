package repository

import (
	"main/config"
	"main/file/model"
)

func Create(file *model.File) {
	config.DbConnection().Create(&file)
}

func GetByFilePath(path string) model.File {
	file := model.File{Path: path}
	config.DbConnection().First(&file)
	return file
}
