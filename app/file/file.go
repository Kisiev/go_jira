package file

import (
	"errors"
	"fmt"
	"main/file/model"
	"math/rand"
	"os"
	"time"
)

func GetRandomFilepath(files []model.File) (model.File, error) {
	fileCount := len(files)

	if fileCount == 0 {
		return model.File{}, fmt.Errorf("нет файлов")
	}

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := fileCount
	randomFileIndex := rand.Intn(max-min) + min

	fullPath := files[randomFileIndex].Path

	if _, err := os.Stat(fullPath); errors.Is(err, os.ErrNotExist) {
		return model.File{}, err
	}

	return files[randomFileIndex], nil
}
