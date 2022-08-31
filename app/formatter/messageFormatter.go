package formatter

import (
	"fmt"
	"main/model"
)

func FormatMessage(task model.Task) string {
	return fmt.Sprintf("📚\t<b>Задача</b>: %s\n📎\t<b>Ссылка</b>: %s\n⚡️\t<b>Приоритет</b>: %s\n⚠️\t<b>Статус</b>: %s\n\n", task.Title, task.Url, priorityFormat(task.Priority), task.Status)
}

func priorityFormat(priority string) string {
	priorityMap := map[string]string{
		"High":        "🟠",
		"Low":         "🔵",
		"Medium":      "🟢",
		"Блокирующий": "⛔️",
	}

	if mapPriority, found := priorityMap[priority]; found {
		return mapPriority
	}

	return "🔴"
}
