package formatter

import (
	"main/entity"
)

type JiraFormatter struct{}

func (f JiraFormatter) Format(task entity.JiraTask) []entity.Task {

	var tasks []entity.Task

	for _, item := range task.Issues {
		task := entity.Task{
			Url:      item.GetUrl(),
			Priority: mapPriority(item.Fields.Priority.Name),
			Title:    item.Fields.Summary,
			Status:   item.Fields.Status.Name,
		}

		tasks = append(tasks, task)
	}

	return tasks
}

func mapPriority(name string) int {
	priorityMap := map[string]int{
		"Highest":     2,
		"High":        1,
		"Low":         -1,
		"Medium":      0,
		"Блокирующий": 3,
	}

	if mapPriority, found := priorityMap[name]; found {
		return mapPriority
	}

	return -2
}
