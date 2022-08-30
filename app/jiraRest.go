package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/entity"
	"main/helper"
	"net/http"
	"strings"
)

type taskRequest struct {
	Jql        string   `json:"jql"`
	StartAt    int      `json:"startAt"`
	MaxResults int      `json:"maxResults"`
	Fields     []string `json:"fields"`
}

func getData() entity.Task {

	var task entity.Task
	var payloadData = taskRequest{
		Jql:        fmt.Sprintf("project = TRACEWAY and assignee=%s and status not in (Закрыто, Выполнено, Done, CLOSED, Canceled) ORDER BY created DESC", helper.GetEnv("JIRA_USER", "")),
		StartAt:    0,
		MaxResults: 20,
		Fields: []string{
			"id",
			"key",
			"summary",
			"status",
			"versions",
			"description",
			"priority",
			"issuetype",
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

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
