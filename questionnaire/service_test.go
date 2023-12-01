package questionnaire

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetQuestionnaire(t *testing.T) {
	ctx := context.Background()

	givenId := "123"
	want := Questionnaire{
		Id:                   givenId,
		StudyId:              "234",
		Name:                 "345",
		Questions:            "456",
		MaxAttempts:          10,
		HoursBetweenAttempts: 23,
	}

	dataClientMock := mockDataClient{
		returnedValue: want,
		returnedError: nil,
	}
	service := Service{
		DataClient: dataClientMock,
	}

	questionnaire, err := service.GetQuestionnaire(ctx, givenId)
	assert.NoError(t, err)
	assert.Equal(t, want, questionnaire)

}

type mockDataClient struct {
	returnedValue Questionnaire
	returnedError error
}

func (m mockDataClient) GetQuestionnaire(ctx context.Context, id string) (Questionnaire, error) {
	return m.returnedValue, m.returnedError
}
