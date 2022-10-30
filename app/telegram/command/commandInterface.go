package command

import "main/telegram/entity"

type Command interface {
	Run(update entity.TelegramUpdate)
	Support(update entity.TelegramUpdate) bool
}

type KeyboardCommand interface {
	Run(update entity.TelegramUpdateInline)
	Support(update entity.TelegramUpdateInline) bool
}
