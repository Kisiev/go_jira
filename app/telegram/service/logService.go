package service

import (
	"encoding/json"
	telegramEntity "main/telegram/entity"
	"main/telegram/model"
	"main/telegram/repository"
)

type LogService struct{}

func (l LogService) LoggingFromUpdateEntity(update telegramEntity.TelegramUpdate) {
	telegramUpdateSrt, err := json.Marshal(update)

	if err != nil {
		return
	}

	log := model.Log{IsBot: false, TelegramId: update.Message.From.Id, Payload: string(telegramUpdateSrt)}
	repository.Create(&log)
}
