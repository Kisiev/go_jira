package formatter

import (
	"main/jira/entity"
)

type FormattedTasks struct {
	Url      string
	Priority string
	Title    string
	Status   string
}

type Formatter interface {
	Format(task entity.JiraTask) []entity.Task
}
