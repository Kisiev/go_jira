package entity

type TelegramUpdate struct {
	UpdateId int `json:"update_id"`
	Message  struct {
		MessageId int `json:"message_id"`
		From      struct {
			Id           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			Id        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date     int    `json:"date"`
		Text     string `json:"text"`
		Entities []struct {
			Offset int    `json:"offset"`
			Length int    `json:"length"`
			Type   string `json:"type"`
		} `json:"entities"`
	} `json:"message"`
}

type TelegramUpdateInline struct {
	UpdateId      int `json:"update_id"`
	CallbackQuery struct {
		Id   string `json:"id"`
		From struct {
			Id           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Message struct {
			MessageId int `json:"message_id"`
			From      struct {
				Id        int    `json:"id"`
				IsBot     bool   `json:"is_bot"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
			} `json:"from"`
			Chat struct {
				Id        int    `json:"id"`
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
				Type      string `json:"type"`
			} `json:"chat"`
			Date        int    `json:"date"`
			Text        string `json:"text"`
			ReplyMarkup struct {
				InlineKeyboard [][]struct {
					Text         string `json:"text"`
					Url          string `json:"url,omitempty"`
					CallbackData string `json:"callback_data,omitempty"`
				} `json:"inline_keyboard"`
			} `json:"reply_markup"`
		} `json:"message"`
		ChatInstance string `json:"chat_instance"`
		Data         string `json:"data"`
	} `json:"callback_query"`
}
