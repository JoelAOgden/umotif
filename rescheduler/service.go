package main

import (
	"context"
	"errors"
	"github.com/JoelAOgden/umotif/questionnaire"
	"github.com/JoelAOgden/umotif/queue"
	"sync"
)

type questionnaireService interface {
	GetQuestionnaire(ctx context.Context, id string) (questionnaire.Questionnaire, error)
}

type scheduleQuestionnaireService interface {
	SubmitNewSchedule(ctx context.Context, questionnaireId string, userId string, completedAt string, hoursBetweenAttempts int) error
}

type queueService interface {
	PushToQueue(ctx context.Context, msg queue.Message) error
}

type reschedulerService struct {
	// Although the interfaces don't prevent coupling here they are very useful for testing
	QuestionnaireService         questionnaireService
	ScheduleQuestionnaireService scheduleQuestionnaireService
	SqsService                   queueService
}

type questionnaireCompletedInput struct {
	userId               string
	questionnaireId      string
	completedAt          string
	remainingCompletions int
}

func (s reschedulerService) SubmitQuestionnaireCompletion(ctx context.Context, input questionnaireCompletedInput) error {

	if !s.NewScheduleRequired(input.remainingCompletions) {
		return s.CompleteQuestionnaire(ctx)
	}

	currentQuestionnaire, err := s.QuestionnaireService.GetQuestionnaire(ctx, input.questionnaireId)
	if err != nil {
		return err
	}

	return s.ScheduleNewQuestionnaire(ctx, input.questionnaireId, input.userId, input.completedAt, currentQuestionnaire.HoursBetweenAttempts)
}

func (s reschedulerService) NewScheduleRequired(RemainingCompletions int) bool {

	// this can be expanded to include other conditions as needed
	// I'm not 100% certain what those conditions are if I'm honest
	// Not sure if remaining completions should be compared to the questionnaire.
	// I've left this mostly empty, but it's pretty trivial to add more if needed

	return RemainingCompletions > 0
}

func (s reschedulerService) CompleteQuestionnaire(ctx context.Context) error {
	return s.SqsService.PushToQueue(ctx, queue.Message{Body: "user has completed all questionnaire"}) // todo: find out the messages
}

func (s reschedulerService) ScheduleNewQuestionnaire(ctx context.Context, questionnaireId string, userId string, completedAt string, HoursBetweenAttempts int) error {

	errorChan := make(chan error)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := s.ScheduleQuestionnaireService.SubmitNewSchedule(ctx, questionnaireId, userId, completedAt, HoursBetweenAttempts)
		if err != nil {
			errorChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := s.SqsService.PushToQueue(ctx, queue.Message{Body: "New Schedule Added"}) // todo: find out the messages
		if err != nil {
			errorChan <- err
		}
	}()

	// wait for the wait group to be complete, then close the channel
	// this might be possible to remove with a buffer (consider testing)
	go func() {
		wg.Wait()
		close(errorChan)
	}()

	// lock on the channel waiting for the goroutines to complete
	// will unlock on channel close abxove
	var returnError error
	for err := range errorChan {
		if returnError == nil {
			returnError = err
			continue
		}
		returnError = errors.Join(returnError, err)
	}

	return returnError

}
