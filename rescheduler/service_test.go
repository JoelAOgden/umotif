package rescheduler

import (
	"context"
	"github.com/JoelAOgden/umotif/questionnaire"
	"github.com/JoelAOgden/umotif/queue"
	"testing"
)

func TestSubmitQuestionnaireCompletion(t *testing.T) {
	// I'm running out of time but essentially, just test using the mocks below
	// an example I "finished" can be seen in the schedulerQuestionnaire package
}

func TestNewScheduleRequired(t *testing.T) {

}

func TestCompleteQuestionnaire(t *testing.T) {

}

func TestScheduleNewQuestionnaire(t *testing.T) {

}

// Mocks

type questionnaireServiceMock struct {
	onGetQuestionnaire func(ctx context.Context, id string) (questionnaire.Questionnaire, error)
}

func (m questionnaireServiceMock) GetQuestionnaire(ctx context.Context, id string) (questionnaire.Questionnaire, error) {
	return m.onGetQuestionnaire(ctx, id)
}

type scheduleQuestionnaireServiceMock struct {
	onSubmitNewSchedule func(ctx context.Context, questionnaireId string, userId string, completedAt string, hoursBetweenAttempts int) error
}

func (m scheduleQuestionnaireServiceMock) SubmitNewSchedule(ctx context.Context, questionnaireId string, userId string, completedAt string, hoursBetweenAttempts int) error {
	return m.onSubmitNewSchedule(ctx, questionnaireId, userId, completedAt, hoursBetweenAttempts)
}

type queueServiceMock struct {
	onPushToQueue func(ctx context.Context, msg queue.Message) error
}

func (m queueServiceMock) PushToQueue(ctx context.Context, msg queue.Message) error {
	return m.onPushToQueue(ctx, msg)
}
