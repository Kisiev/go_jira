package command

import (
	"fmt"
	"main/jira"
	"main/jira/repository"
	"main/telegram"
	telegramEntity "main/telegram/entity"
	"sort"
	"strconv"
	"time"
)

type JiraWorkLogCommand struct{}

type workLogData struct {
	Url       string
	SpendTime float64
	Day       string
	Date      string
}

var userWorkLog map[int64][]workLogData

func (j JiraWorkLogCommand) Run(update telegramEntity.TelegramUpdate) {
	var bot telegram.BotInterface = telegram.Bot{}

	telegramMessage := update.Message
	user := repository.FindJiraUserByTelegramId(telegramMessage.From.Id)

	dateStart := time.Now().AddDate(0, 0, -10).Format("2006-01-02")
	dateEnd := time.Now().AddDate(0, 0, 2).Format("2006-01-02")

	rawJiraFilter := fmt.Sprintf("worklogDate >= '%s' and worklogDate < '%s' and worklogAuthor = %s ORDER BY priority DESC, created DESC", dateStart, dateEnd, user.UserName)
	tasks := jira.GetTasksForUser(rawJiraFilter)

	userWorkLog = make(map[int64][]workLogData)

	for _, task := range tasks.Issues {
		workLog := jira.GetTaskWorkLog(task.Key)
		for _, workLog := range workLog.WorkLogs {
			if workLog.Author.Name != user.UserName {
				continue
			}

			workDate, err := time.Parse("2006-01-02T15:04:05-0700", workLog.Started)
			if err != nil {
				continue
			}

			date, _ := time.Parse("2006-01-02", dateStart)
			if workDate.Before(date) {
				continue
			}

			key, err := time.Parse("2006-01-02", workDate.Format("2006-01-02"))
			if err != nil {
				continue
			}

			newElement := workLogData{
				Url:       task.GetUrl(),
				SpendTime: float64(workLog.TimeSpentSeconds) / 60 / 60,
				Day:       dayMap(workDate.Weekday().String()),
				Date:      workDate.Format("2006-01-02"),
			}

			userWorkLog[key.Unix()] = append(userWorkLog[key.Unix()], newElement)
		}
	}
	formatMessage := formatMessage(userWorkLog)
	bot.SimpleSendMessage(formatMessage, strconv.Itoa(user.User.TelegramId))
}

func (j JiraWorkLogCommand) Support(update telegramEntity.TelegramUpdate) bool {
	return true
}

func formatMessage(userWorkLog map[int64][]workLogData) string {
	var message string
	var keys []int

	for key := range userWorkLog {
		keys = append(keys, int(key))
	}
	sort.Ints(keys)

	for _, key := range keys {
		tasks := userWorkLog[int64(key)]
		message += tasks[0].Day + " " + tasks[0].Date + "\n\t\t"
		totalDayHours := 0.0

		for _, workLogItem := range tasks {
			totalDayHours += workLogItem.SpendTime
			message += fmt.Sprintf("%s - %.1f\n\t\t", workLogItem.Url, workLogItem.SpendTime)
		}
		message += fmt.Sprintf("Итого: %.1f\n\n", totalDayHours)
	}
	return message
}

func dayMap(day string) string {
	days := map[string]string{
		"Sunday":    "Вс",
		"Monday":    "Пн",
		"Tuesday":   "Вт",
		"Wednesday": "Ср",
		"Thursday":  "Чт",
		"Friday":    "Пт",
		"Saturday":  "Сб",
	}

	found, _ := days[day]
	return found
}
