package command

import (
	"encoding/json"
	jiraCommand "main/jira/command"
	"main/telegram/constant"
	"main/telegram/entity"
	"main/user/model"
	"main/user/repository"
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
}

func tryHandle(command string, update entity.TelegramUpdate) bool {
	var commandMap = map[string]Command{
		constant.ActionStartCommand:    StartCommand{},
		constant.ActionSetJiraUserName: jiraCommand.UserCreateCommand{},
		constant.ActionTaskView:        jiraCommand.JiraTaskCommand{},
		constant.ActionReport:          jiraCommand.JiraReportCommand{},
		constant.ActionWorkLog:         jiraCommand.JiraWorkLogCommand{},
	}

	if commandMap, found := commandMap[command]; found {
		commandMap.Run(update)
		return true
	}

	return false
}
