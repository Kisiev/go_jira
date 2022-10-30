package keyboardCommand

import (
	tgCommand "main/telegram/command"
	"main/telegram/entity"
	userCommand "main/user/command"
	"main/user/constant"
	"regexp"
)

func Handle(update entity.TelegramUpdateInline) {
	if tryHandle(update.CallbackQuery.Data, update) {
		return
	}

	if tryHandleWithRegex(update.CallbackQuery.Data, update) {
		return
	}
}

func tryHandle(command string, update entity.TelegramUpdateInline) bool {
	var commandMap = map[string]tgCommand.KeyboardCommand{
		constant.NotificationSetting: userCommand.SettingCommand{},
	}

	if commandMap, found := commandMap[command]; found {
		if commandMap.Support(update) {
			commandMap.Run(update)
		}
		return true
	}

	return false
}

func tryHandleWithRegex(command string, update entity.TelegramUpdateInline) bool {
	var commandMap = map[string]tgCommand.KeyboardCommand{
		constant.NotificationSetting: userCommand.NotificationSwitcherCommand{},
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
