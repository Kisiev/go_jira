package formatter

import (
	"main/jira/entity"
	"main/jira/model"
)

type JiraFormatter struct{}

func (f JiraFormatter) Format(task entity.JiraTask) []model.Task {

	var tasks []model.Task

	for _, item := range task.Issues {
		task := model.Task{
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
