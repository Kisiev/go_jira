package formatter

import (
	"main/entity"
)

type JiraFormatter struct{}

func (f JiraFormatter) Format(task entity.Task) []FormattedTasks {

	var tasks []FormattedTasks

	for _, item := range task.Issues {
		task := FormattedTasks{
			Url:      item.GetUrl(),
			Priority: item.Fields.Priority.Name,
			Title:    item.Fields.Summary,
			Status:   item.Fields.Status.Name,
		}

		tasks = append(tasks, task)
	}

	return tasks
}
