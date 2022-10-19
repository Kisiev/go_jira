package command

import (
	"fmt"
	"main/jira"
	"main/jira/entity"
	"main/jira/repository"
	"main/telegram"
	telegramEntity "main/telegram/entity"
	"strconv"
	"time"
)

type JiraReportCommand struct{}

var yesterdayChan chan entity.JiraTask
var todayChan chan entity.JiraTask

func (j JiraReportCommand) Run(update telegramEntity.TelegramUpdate) {
	var bot telegram.BotInterface = telegram.Bot{}

	telegramMessage := update.Message
	user := repository.FindJiraUserByTelegramId(telegramMessage.From.Id)

	yesterdayChan = make(chan entity.JiraTask)
	dateStart, dateEnd := getPrevDate()
	rawJiraFilter := fmt.Sprintf("worklogDate >= '%s' and worklogDate < '%s' and worklogAuthor = %s ORDER BY priority DESC, created DESC", dateStart, dateEnd, user.UserName)
	go getYesterdayTasks(rawJiraFilter, yesterdayChan)

	todayChan = make(chan entity.JiraTask)
	rawJiraFilter = fmt.Sprintf("assignee=%s and status not in (Закрыто, Выполнено, Done, CLOSED, Canceled) ORDER BY priority DESC, created DESC", user.UserName)
	go getTodayTasks(rawJiraFilter, todayChan)

	message := formatReport(<-yesterdayChan, <-todayChan)

	go bot.SimpleSendMessage(message, strconv.Itoa(telegramMessage.From.Id))
}

func (j JiraReportCommand) Support(update telegramEntity.TelegramUpdate) bool {
	return true
}

func getYesterdayTasks(rawFilter string, yesterdayChan chan entity.JiraTask) {
	yesterdayChan <- jira.GetTasksForUser(rawFilter)
}

func getTodayTasks(rawFilter string, todayChan chan entity.JiraTask) {
	todayChan <- jira.GetTasksForUser(rawFilter)
}

func formatReport(yesterdayTasks entity.JiraTask, todayTasks entity.JiraTask) string {
	message := "Вчера\n"

	if time.Now().Weekday().String() == "Monday" {
		message = "Пятница\n"
	}

	for _, task := range yesterdayTasks.Issues {
		message += fmt.Sprintf("%s\n", task.GetUrl())
	}

	message += "Сегодня\n"

	for index, task := range todayTasks.Issues {
		message += fmt.Sprintf("%s\n", task.GetUrl())

		if index > 2 {
			break
		}
	}

	return message
}

func getPrevDate() (string, string) {
	var dateStart string
	var day = time.Now().Weekday().String()

	dayMap := map[string]int{
		"Monday": 3,
		"Sunday": 2,
	}

	prevDay := 1

	if dayMap, found := dayMap[day]; found {
		prevDay = dayMap
	}

	dateStart = time.Now().AddDate(0, 0, -prevDay).Format("2006-01-02")
	endStart := time.Now().Format("2006-01-02")

	return dateStart, endStart
}
