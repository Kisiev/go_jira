package command

import "main/telegram/entity"

type Command interface {
	Run(update entity.TelegramUpdate)
}
