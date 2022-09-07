package entity

import "main/helper"

type JiraTask struct {
	Expand     string   `json:"expand"`
	StartAt    int      `json:"startAt"`
	MaxResults int      `json:"maxResults"`
	Total      int      `json:"total"`
	Issues     []Issues `json:"issues"`
}

type Issues struct {
	Expand string `json:"expand"`
	Id     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields struct {
		Summary   string `json:"summary"`
		Issuetype struct {
			Self        string `json:"self"`
			Id          string `json:"id"`
			Description string `json:"description"`
			IconUrl     string `json:"iconUrl"`
			Name        string `json:"name"`
			Subtask     bool   `json:"subtask"`
			AvatarId    int    `json:"avatarId"`
		} `json:"issuetype"`
		Versions []struct {
			Self        string `json:"self"`
			Id          string `json:"id"`
			Name        string `json:"name"`
			Archived    bool   `json:"archived"`
			Released    bool   `json:"released"`
			Description string `json:"description,omitempty"`
			ReleaseDate string `json:"releaseDate,omitempty"`
		} `json:"versions"`
		TimeTracking struct {
			RemainingEstimate        string `json:"remainingEstimate"`
			TimeSpent                string `json:"timeSpent"`
			RemainingEstimateSeconds int    `json:"remainingEstimateSeconds"`
			TimeSpentSeconds         int    `json:"timeSpentSeconds"`
		} `json:"timetracking"`
		Description string `json:"description"`
		Priority    struct {
			Self    string `json:"self"`
			IconUrl string `json:"iconUrl"`
			Name    string `json:"name"`
			Id      string `json:"id"`
		} `json:"priority"`
		Status struct {
			Self           string `json:"self"`
			Description    string `json:"description"`
			IconUrl        string `json:"iconUrl"`
			Name           string `json:"name"`
			Id             string `json:"id"`
			StatusCategory struct {
				Self      string `json:"self"`
				Id        int    `json:"id"`
				Key       string `json:"key"`
				ColorName string `json:"colorName"`
				Name      string `json:"name"`
			} `json:"statusCategory"`
		} `json:"status"`
	} `json:"fields"`
}

func (i *Issues) GetUrl() string {
	return helper.GetEnv("JIRA_URL", "") + "/browse/" + i.Key
}
