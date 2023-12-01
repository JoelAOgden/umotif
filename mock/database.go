package mock

import (
	"context"
	"github.com/JoelAOgden/umotif/questionnaire"
	"github.com/JoelAOgden/umotif/scheduleQuestionnaire"
)

type DataBaseClient struct {
	// todo: this
	questionnaires        map[string]questionnaire.Questionnaire
	scheduleQuestionnaire map[string]scheduleQuestionnaire.ScheduledQuestionnaire
}

func NewMockDatabaseClient() *DataBaseClient {
	return &DataBaseClient{
		questionnaires:        make(map[string]questionnaire.Questionnaire),
		scheduleQuestionnaire: make(map[string]scheduleQuestionnaire.ScheduledQuestionnaire),
	}
}

func (d DataBaseClient) GetQuestionnaire(ctx context.Context, id string) (questionnaire.Questionnaire, error) {
	// todo: add error handling
	return d.questionnaires[id], nil
}

func (d DataBaseClient) PutScheduleQuestionnaire(ctx context.Context, schedule scheduleQuestionnaire.ScheduledQuestionnaire) error {
	d.scheduleQuestionnaire[schedule.Id] = schedule
	// todo: add error handling
	return nil
}
