package formatter

import (
	"main/jira/entity"
	"main/jira/model"
)

type FormattedTasks struct {
	Url      string
	Priority string
	Title    string
	Status   string
}

type Formatter interface {
	Format(task entity.JiraTask) []model.Task
}
