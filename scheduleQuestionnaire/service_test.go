package scheduleQuestionnaire

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubmitNewSchedule(t *testing.T) {

	// these tests need to be expanded quite a bit

	questionnaireId := "234"
	userId := "345"
	completedAt := "2020-01-02T15:00:00"
	hoursBetween := 3

	// simple mock function to check the put request
	onClientPut := func(ctx context.Context, schedule ScheduledQuestionnaire) error {
		assert.Equal(t, schedule.QuestionnaireId, questionnaireId)
		assert.Equal(t, schedule.ParticipantId, userId)
		assert.Equal(t, schedule.ScheduledAt, "2020-01-02T18:00:00")
		assert.Equal(t, schedule.Status, StatusPending)
		return nil
	}

	mockClient := clientMock{
		onPut: onClientPut,
	}

	service := Service{
		DataClient: mockClient,
	}

	err := service.SubmitNewSchedule(context.Background(), questionnaireId, userId, completedAt, hoursBetween)
	assert.NoError(t, err)
}

type clientMock struct {
	onPut func(ctx context.Context, schedule ScheduledQuestionnaire) error
}

func (m clientMock) Put(ctx context.Context, schedule ScheduledQuestionnaire) error {
	return m.onPut(ctx, schedule)
}
