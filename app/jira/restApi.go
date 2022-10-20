package jira

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"main/helper"
	"main/jira/entity"
	"net/http"
	"strings"
)

type taskRequest struct {
	Jql        string   `json:"jql"`
	StartAt    int      `json:"startAt"`
	MaxResults int      `json:"maxResults"`
	Fields     []string `json:"fields"`
}

func GetTasksForUser(filter string) entity.JiraTask {

	var task entity.JiraTask
	var payloadData = taskRequest{
		Jql:        filter,
		StartAt:    0,
		MaxResults: 200,
		Fields: []string{
			"id",
			"key",
			"summary",
			"status",
			"versions",
			"description",
			"priority",
			"issuetype",
			"timetracking",
		},
	}

	url := helper.GetEnv("JIRA_URL", "") + "/rest/api/2/search"
	method := "POST"

	encodePayload, err := json.Marshal(payloadData)
	if err != nil {
		log.Fatal(err)
	}

	payload := strings.NewReader(string(encodePayload))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth(helper.GetEnv("JIRA_USER", ""), helper.GetEnv("JIRA_PASS", "")))

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		log.Fatal(err)
	}

	return task
}

func GetTaskWorkLog(issueIdOrKey string) entity.WorkLog {
	var workLog entity.WorkLog

	url := helper.GetEnv("JIRA_URL", "") + "/rest/api/2/issue/" + issueIdOrKey + "/worklog"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth(helper.GetEnv("JIRA_USER", ""), helper.GetEnv("JIRA_PASS", "")))

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &workLog)
	if err != nil {
		log.Fatal(err)
	}

	return workLog
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
