package main

import (
	"context"
	"github.com/JoelAOgden/umotif/mock"
	"github.com/JoelAOgden/umotif/questionnaire"
	"github.com/JoelAOgden/umotif/queue"
	"github.com/JoelAOgden/umotif/scheduleQuestionnaire"
	"github.com/aws/aws-lambda-go/lambda"
)

// The incoming event
type QuestionnaireCompletedEvent struct {
	Id                   string
	UserId               string
	StudyId              string
	QuestionnaireId      string
	CompletedAt          string
	RemainingCompletions int
}

func LambdaHandler(ctx context.Context, event QuestionnaireCompletedEvent) (string, error) {

	mockDatabaseClient := mock.NewMockDatabaseClient()
	mockSqsClient := mock.SqsClient{}

	// todo: this
	rescheduler := reschedulerService{
		questionnaire.Service{
			DataClient: mockDatabaseClient,
		},
		scheduleQuestionnaire.Service{
			DataClient: mockDatabaseClient,
		},
		queue.SqsService{
			Client: mockSqsClient,
		},
	}

	err := rescheduler.SubmitQuestionnaireCompletion(ctx, questionnaireCompletedInput{
		userId:               event.UserId,
		questionnaireId:      event.QuestionnaireId,
		completedAt:          event.CompletedAt,
		remainingCompletions: event.RemainingCompletions,
	})
	if err != nil {
		return "", err
	}

	return "", nil
}

func main() {
	lambda.Start(LambdaHandler)
}
