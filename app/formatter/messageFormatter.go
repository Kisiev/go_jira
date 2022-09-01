package formatter

import (
	"fmt"
	"main/entity"
)

func FormatMessage(task entity.Task) string {
	return fmt.Sprintf("📚\t<b>Задача</b>: %s\n📎\t<b>Ссылка</b>: %s\n⚡️\t<b>Приоритет</b>: %s\n⚠️\t<b>Статус</b>: %s\n\n", task.Title, task.Url, priorityFormat(task.Priority), task.Status)
}

func priorityFormat(priority int) string {
	priorityMap := map[int]string{
		2:  "🔴",
		1:  "🟠",
		-1: "🔵",
		0:  "🟢",
		3:  "⛔️",
	}

	if mapPriority, found := priorityMap[priority]; found {
		return mapPriority
	}

	return "🔵"
}
