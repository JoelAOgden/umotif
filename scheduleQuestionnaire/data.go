package scheduleQuestionnaire

// CREATE TABLE scheduled_questionnaires (
//    id VARCHAR(128) PRIMARY KEY NOT NULL,
//    questionnaire_id VARCHAR(128) NOT NULL,
//    participant_id VARCHAR(128) NOT NULL,
//    scheduled_at DATETIME NOT NULL,
//    status ENUM(pending,completed)
//);

type ScheduledQuestionnaire struct {
	// this will need to be updated when not using mocks
	Id              string
	QuestionnaireId string
	ParticipantId   string
	ScheduledAt     string
	Status          string
}
