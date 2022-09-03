package command

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/helper"
	"main/jira/entity"
	taskFormatter "main/jira/formatter"
	"main/jira/repository"
	"main/telegram"
	telegramEntity "main/telegram/entity"
	"net/http"
	"strconv"
	"strings"
)

type taskRequest struct {
	Jql        string   `json:"jql"`
	StartAt    int      `json:"startAt"`
	MaxResults int      `json:"maxResults"`
	Fields     []string `json:"fields"`
}

type JiraTaskCommand struct{}

func GetTasksForUser(userName string) entity.JiraTask {

	var task entity.JiraTask
	var payloadData = taskRequest{
		Jql:        fmt.Sprintf("project = TRACEWAY and assignee=%s and status not in (Закрыто, Выполнено, Done, CLOSED, Canceled) ORDER BY created DESC", userName),
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

func (u JiraTaskCommand) Run(update telegramEntity.TelegramUpdate) {
	var bot telegram.BotInterface = telegram.Bot{}

	telegramMessage := update.Message
	user := repository.FindJiraUserByTelegramId(telegramMessage.From.Id)

	jiraData := GetTasksForUser(user.UserName)

	var formatter taskFormatter.Formatter = taskFormatter.JiraFormatter{}
	data := formatter.Format(jiraData)

	var message string

	for _, item := range data {
		message += taskFormatter.FormatMessage(item)
	}

	if len(message) == 0 {
		message = fmt.Sprintf("Для пользователя %s не найдены задачи", user.UserName)
	}

	go bot.SimpleSendMessage(message, strconv.Itoa(telegramMessage.From.Id))
}
