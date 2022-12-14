package formatter

import (
	"fmt"
	"main/jira/model"
)

func FormatMessage(task model.Task) string {
	return fmt.Sprintf("📚\t<b>Задача</b>: %s\n📎\t<b>Ссылка</b>: %s\n⚡️\t<b>Приоритет</b>: %s\n⚠️\t<b>Статус</b>: %s\n🏷\t<b>Тип</b>: %s\n\n", task.Title, task.Url, priorityFormat(task.Priority), task.Status, task.Type)
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
