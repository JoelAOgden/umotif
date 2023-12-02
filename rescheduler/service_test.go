package rescheduler

import (
	"context"
	"github.com/JoelAOgden/umotif/questionnaire"
	"github.com/JoelAOgden/umotif/queue"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubmitQuestionnaireCompletion_Required(t *testing.T) {

	ctx := context.Background()

	givenUserId := "123"
	givenQuestionnaireId := "234"
	givenCompletedAt := "345"
	givenRemainingCompletions := 1

	givenQuestionnaire := questionnaire.Questionnaire{
		Id:                   givenQuestionnaireId,
		StudyId:              "studyId",
		Name:                 "name",
		Questions:            "questions",
		MaxAttempts:          7,
		HoursBetweenAttempts: 2,
	}
	wantQuestionnaireId := givenQuestionnaireId
	mockQuestionnaireService := questionnaireServiceMock{
		onGetQuestionnaire: func(ctx context.Context, id string) (questionnaire.Questionnaire, error) {
			assert.Equal(t, id, wantQuestionnaireId)
			return givenQuestionnaire, nil
		},
	}

	scheduleQuestionnaireServiceTriggered := false
	mockScheduleQuestionnaireService := scheduleQuestionnaireServiceMock{
		onSubmitNewSchedule: func(ctx context.Context, questionnaireId string, userId string, completedAt string, hoursBetweenAttempts int) error {
			scheduleQuestionnaireServiceTriggered = true
			assert.Equal(t, questionnaireId, givenQuestionnaireId)
			assert.Equal(t, userId, givenUserId)
			assert.Equal(t, completedAt, givenCompletedAt)
			assert.Equal(t, hoursBetweenAttempts, givenQuestionnaire.HoursBetweenAttempts)
			return nil
		},
	}

	sqsWant := queue.Message{"New Schedule Added"}

	sqsTriggered := false
	mockSqs := queueServiceMock{
		onPushToQueue: func(ctx context.Context, msg queue.Message) error {
			assert.Equal(t, sqsWant, msg)
			sqsTriggered = true
			return nil
		},
	}
	service := Service{
		QuestionnaireService:         mockQuestionnaireService,
		ScheduleQuestionnaireService: mockScheduleQuestionnaireService,
		SqsService:                   mockSqs,
	}

	err := service.SubmitQuestionnaireCompletion(ctx, QuestionnaireCompletedInput{
		UserId:               givenUserId,
		QuestionnaireId:      givenQuestionnaireId,
		CompletedAt:          givenCompletedAt,
		RemainingCompletions: givenRemainingCompletions,
	})

	assert.NoError(t, err)
	assert.True(t, scheduleQuestionnaireServiceTriggered)
	assert.True(t, sqsTriggered)

}

func TestSubmitQuestionnaireCompletion_NotRequired(t *testing.T) {

	ctx := context.Background()

	givenUserId := "123"
	givenQuestionnaireId := "234"
	givenCompletedAt := "345"
	givenRemainingCompletions := 0

	want := queue.Message{"user has completed all questionnaire"}

	sqsTriggered := false
	mockSqs := queueServiceMock{
		onPushToQueue: func(ctx context.Context, msg queue.Message) error {
			assert.Equal(t, want, msg)
			sqsTriggered = true
			return nil
		},
	}
	service := Service{
		SqsService: mockSqs,
	}

	err := service.SubmitQuestionnaireCompletion(ctx, QuestionnaireCompletedInput{
		UserId:               givenUserId,
		QuestionnaireId:      givenQuestionnaireId,
		CompletedAt:          givenCompletedAt,
		RemainingCompletions: givenRemainingCompletions,
	})

	assert.NoError(t, err)
	assert.True(t, sqsTriggered)

}

func TestNewScheduleRequired(t *testing.T) {

	given := 1
	got := newScheduleRequired(given)
	assert.True(t, got)

	given = 0
	got = newScheduleRequired(given)
	assert.False(t, got)

	given = -1
	got = newScheduleRequired(given)
	assert.False(t, got)

}

func TestCompleteQuestionnaire(t *testing.T) {

	ctx := context.Background()
	want := queue.Message{"user has completed all questionnaire"}

	sqsTriggered := false
	mockSqs := queueServiceMock{
		onPushToQueue: func(ctx context.Context, msg queue.Message) error {
			assert.Equal(t, want, msg)
			sqsTriggered = true
			return nil
		},
	}

	service := Service{
		SqsService: mockSqs,
	}

	err := service.CompleteQuestionnaire(ctx)
	assert.NoError(t, err)
	assert.True(t, sqsTriggered)
}

func TestScheduleNewQuestionnaire(t *testing.T) {
	ctx := context.Background()

	givenQuestionnaireId := "123"
	givenUserId := "234"
	givenCompletedAt := "2020-01-02T15:00:00"
	givenHoursBetweenAttempts := 4

	want := queue.Message{"New Schedule Added"}

	sqsTriggered := false
	mockSqs := queueServiceMock{
		onPushToQueue: func(ctx context.Context, msg queue.Message) error {
			assert.Equal(t, want, msg)
			sqsTriggered = true
			return nil
		},
	}

	newScheduleSubmitted := false
	mockScheduleQuestionnaireService := scheduleQuestionnaireServiceMock{
		onSubmitNewSchedule: func(ctx context.Context, questionnaireId string, userId string, completedAt string, hoursBetweenAttempts int) error {
			newScheduleSubmitted = true
			assert.Equal(t, givenQuestionnaireId, questionnaireId)
			assert.Equal(t, givenQuestionnaireId, questionnaireId)
			assert.Equal(t, givenUserId, userId)
			assert.Equal(t, givenCompletedAt, completedAt)
			assert.Equal(t, givenHoursBetweenAttempts, hoursBetweenAttempts)

			return nil
		},
	}

	service := Service{
		ScheduleQuestionnaireService: mockScheduleQuestionnaireService,
		SqsService:                   mockSqs,
	}

	err := service.ScheduleNewQuestionnaire(ctx, givenQuestionnaireId, givenUserId, givenCompletedAt, givenHoursBetweenAttempts)
	assert.NoError(t, err)
	assert.True(t, sqsTriggered)
	assert.True(t, newScheduleSubmitted)

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
