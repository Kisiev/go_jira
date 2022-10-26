package command

import (
	"encoding/json"
	fileCommand "main/file/command"
	jiraCommand "main/jira/command"
	"main/telegram/constant"
	"main/telegram/entity"
	"main/user/model"
	"main/user/repository"
	"regexp"
)

func Handle(update entity.TelegramUpdate) {
	if tryHandle(update.Message.Text, update) {
		return
	}

	user := repository.FindByTelegramId(update.Message.From.Id)
	var nextAction model.NextAction

	err := json.Unmarshal([]byte(user.NextAction), &nextAction)
	if err != nil {
		return
	}

	if tryHandle(nextAction.Action, update) {
		return
	}

	if tryHandleWithRegex(update.Message.Text, update) {
		return
	}

	sendUnknownCommand(update)
}

func sendUnknownCommand(update entity.TelegramUpdate) {
	unknownCommand := UnknownCommand{}
	unknownCommand.Run(update)
}

func tryHandleWithRegex(command string, update entity.TelegramUpdate) bool {
	var commandMap = map[string]Command{
		constant.ActionLinkToPicture: fileCommand.UploadCommand{},
	}

	for regEx, handler := range commandMap {
		compile, err := regexp.Compile(regEx)
		if err != nil {
			return false
		}

		if compile.MatchString(command) && handler.Support(update) {
			handler.Run(update)
		}
	}

	return false
}

func tryHandle(command string, update entity.TelegramUpdate) bool {
	var commandMap = map[string]Command{
		constant.ActionStartCommand:    StartCommand{},
		constant.ActionSetJiraUserName: jiraCommand.UserCreateCommand{},
		constant.ActionTaskView:        jiraCommand.JiraTaskCommand{},
		constant.ActionReport:          jiraCommand.JiraReportCommand{},
		constant.ActionWorkLog:         jiraCommand.JiraWorkLogCommand{},
		constant.ActionRandPicture:     RandomFileCommand{},
	}

	if commandMap, found := commandMap[command]; found {
		if commandMap.Support(update) {
			commandMap.Run(update)
		}
		return true
	}

	return false
}
