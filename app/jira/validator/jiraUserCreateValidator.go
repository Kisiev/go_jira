package validator

import (
	"errors"
	"main/user/model"
	"regexp"
)

func Validate(user model.JiraUser) error {
	_, err := validateUserName(user.UserName)
	if err != nil {
		return err
	}
	return nil
}

func validateUserName(userName string) (string, error) {
	reg, _ := regexp.Compile(`[a-z]{1,2}\.[a-z]+`)

	valid := reg.FindString(userName)

	if len(valid) == 0 {
		return "", errors.New("некорректное имя пользователя для JIRA")
	}

	return "", nil
}
