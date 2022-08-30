package helper

import (
	"os"
)

func GetEnv(key, defaultValue string) string {
	value, err := os.LookupEnv(key)

	if !err {
		return defaultValue
	}

	return value
}
