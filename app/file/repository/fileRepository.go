package repository

import (
	"main/config"
	"main/file/entity"
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

func GetById(id int) model.File {
	file := model.File{}
	config.DbConnection().Where("id = ?", id).First(&file)
	return file
}

func GetRandPathForUser(userId int) entity.FileCount {
	var fileCount entity.FileCount
	config.DbConnection().Raw("select files.id as id, coalesce(count, 0) as count "+
		"From files "+
		"left join file_loggings fl on files.id = fl.file_id and user_id = ?"+
		"order by count, random() limit 1", userId).First(&fileCount)

	return fileCount
}

func AddCountToFileForUser(fileId, userId, count int) {
	if config.DbConnection().Table("file_loggings").
		Where("file_id = ? and user_id = ?", fileId, userId).
		Update("file_id", fileId).
		Update("user_id", userId).
		Update("count", count).
		RowsAffected == 0 {
		config.DbConnection().Create(&model.FileLogging{FileID: fileId, UserID: userId, Count: count})
	}
}
