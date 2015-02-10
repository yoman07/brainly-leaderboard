package leaderboard

import "testing"

var service *AnswerService

func init() {
	repository := GetRepositoryMock()
	connector := GetBrainlyConnectorMock()
	service = CreateAnswerService(repository, connector)
}

func TestAddAnswerForUser(t *testing.T) {
	created := 1423088988
	userId := 5
	answer, err := service.CreateAnswerForUser(userId, created)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if answer.Created != created {
		t.Errorf("Unexpected created timestamp: %d", answer.Created)
	}

	if answer.UserId != userId {
		t.Errorf("Unexpected userId: %d", answer.UserId)
	}
}
