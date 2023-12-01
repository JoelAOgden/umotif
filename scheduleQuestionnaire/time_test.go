package scheduleQuestionnaire

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateScheduledAtTime(t *testing.T) {

	givenStartTime := "2020-01-02T15:00:00"

	got, err := generateScheduledAtTime(givenStartTime, 3)
	assert.NoError(t, err)

	want := "2020-01-02T18:00:00"

	assert.Equal(t, want, got)
}
