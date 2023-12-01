package main

import (
	"context"
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

	// todo: this
	rescheduler := reschedulerService{
		questionnaire.Service{
			DataClient: nil,
		},
		scheduleQuestionnaire.Service{
			DataClient: nil,
		},
		queue.SqsService{
			Client: nil,
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
