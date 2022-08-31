package formatter

import (
	"main/entity"
	"main/model"
)

type JiraFormatter struct{}

func (f JiraFormatter) Format(task entity.Task) []model.Task {

	var tasks []model.Task

	for _, item := range task.Issues {
		task := model.Task{
			Url:      item.GetUrl(),
			Priority: item.Fields.Priority.Name,
			Title:    item.Fields.Summary,
			Status:   item.Fields.Status.Name,
		}

		tasks = append(tasks, task)
	}

	return tasks
}
