package formatter

import (
	"main/entity"
	"main/model"
)

type FormattedTasks struct {
	Url      string
	Priority string
	Title    string
	Status   string
}

type Formatter interface {
	Format(task entity.Task) []model.Task
}
