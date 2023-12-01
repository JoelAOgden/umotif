package queue

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushToQueue(t *testing.T) {
	ctx := context.Background()

	mockSqsClient := &sqsClientMock{""}
	service := SqsService{
		Client: mockSqsClient,
	}

	msg := "this is a message"
	given := Message{Body: msg}

	err := service.PushToQueue(ctx, given)
	assert.NoError(t, err)
	assert.Equal(t, msg, mockSqsClient.messagePushed)

}

type sqsClientMock struct {
	messagePushed string
}

func (s *sqsClientMock) Push(ctx context.Context, message string) error {
	s.messagePushed = message
	return nil
}
