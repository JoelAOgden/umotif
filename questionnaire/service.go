package questionnaire

import "context"

type questionnairesDataClient interface {
	GetQuestionnaire(ctx context.Context, id string) (Questionnaire, error)
}

type Service struct {
	DataClient questionnairesDataClient
}

func (s Service) GetQuestionnaire(ctx context.Context, id string) (Questionnaire, error) {
	return s.DataClient.GetQuestionnaire(ctx, id)
}
