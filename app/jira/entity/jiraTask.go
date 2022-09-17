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
		WorkLog WorkLog `json:"worklog"`
	} `json:"fields"`
}

type WorkLog struct {
	StartAt    int `json:"startAt"`
	MaxResults int `json:"maxResults"`
	Total      int `json:"total"`
	WorkLogs   []struct {
		Self   string `json:"self"`
		Author struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
		} `json:"author"`
		UpdateAuthor struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
		} `json:"updateAuthor"`
		Comment          string `json:"comment"`
		Created          string `json:"created"`
		Updated          string `json:"updated"`
		Started          string `json:"started"`
		TimeSpent        string `json:"timeSpent"`
		TimeSpentSeconds int    `json:"timeSpentSeconds"`
		ID               string `json:"id"`
		IssueID          string `json:"issueId"`
	} `json:"worklogs"`
}

func (i *Issues) GetUrl() string {
	return helper.GetEnv("JIRA_URL", "") + "/browse/" + i.Key
}
