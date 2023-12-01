package scheduleQuestionnaire

import "time"

const timeFormat = "2006-01-02T15:04:05"

func generateScheduledAtTime(startTime string, hoursBetween int) (string, error) {
	parsedQuestionnaireTime, err := time.Parse(timeFormat, startTime)
	if err != nil {
		return "", err
	}
	scheduledAt := parsedQuestionnaireTime.Add(time.Hour * time.Duration(hoursBetween)).Format(timeFormat)

	return scheduledAt, nil
}
