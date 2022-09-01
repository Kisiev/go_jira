package entity

type Task struct {
	Title    string `json:"title"`
	Url      string `json:"url"`
	Priority int    `json:"priority"`
	Status   string `json:"status"`
}
