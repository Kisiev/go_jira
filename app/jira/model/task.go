package model

type Task struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement:true;column=id"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Priority int    `json:"priority"`
	Status   string `json:"status"`
	UserId   int    `json:"user_id"`
}
