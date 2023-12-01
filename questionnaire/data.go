package questionnaire

type Questionnaire struct {
	// This will need to be updated when not using mocks
	Id                   string
	StudyId              string
	Name                 string
	Questions            string
	MaxAttempts          int
	HoursBetweenAttempts int
}
