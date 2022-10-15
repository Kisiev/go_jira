package file

import (
	"errors"
	"io/ioutil"
	"main/helper"
	"math/rand"
	"os"
	"time"
)

func GetRandomFilepath() (string, error) {
	files, err := ioutil.ReadDir(helper.GetEnv("FILES_PATH", ""))
	if err != nil {
		return "", err
	}

	fileCount := len(files)

	if fileCount == 0 {
		return "", err
	}

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := fileCount
	randomFile := rand.Intn(max-min) + min

	fullPath := helper.GetEnv("FILES_PATH", "") + files[randomFile].Name()

	if _, err = os.Stat(fullPath); errors.Is(err, os.ErrNotExist) {
		return "", err
	}

	return fullPath, nil
}
