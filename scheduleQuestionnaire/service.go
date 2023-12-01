package scheduleQuestionnaire

import (
	"context"
	"github.com/google/uuid"
)

type ScheduledQuestionnairesDataClient interface {
	PutScheduleQuestionnaire(ctx context.Context, schedule ScheduledQuestionnaire) error
}

type Service struct {
	DataClient ScheduledQuestionnairesDataClient
}

func (s Service) SubmitNewSchedule(ctx context.Context, questionnaireId string, userId string, completedAt string, hoursBetweenAttempts int) error {

	scheduledAt, err := generateScheduledAtTime(completedAt, hoursBetweenAttempts)
	if err != nil {
		return err
	}

	newSchedule := ScheduledQuestionnaire{
		Id:              uuid.New().String(),
		QuestionnaireId: questionnaireId,
		ParticipantId:   userId,
		ScheduledAt:     scheduledAt,
		Status:          StatusPending,
	}

	err = s.DataClient.PutScheduleQuestionnaire(ctx, newSchedule)
	if err != nil {
		return err
	}

	return nil
}
