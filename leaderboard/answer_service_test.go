package leaderboard

import "testing"

var answerService *AnswerService

func answerServiceTest() {
    repo := GetRepositoryMock()

    repo.SaveUser(
        User{ProfileUrl: "http://user/1"},
        )
    repo.SaveUser(
        User{ProfileUrl: "http://user/2"},
        )

    conn := GetBrainlyConnectorMock()

    conn.AddAnswer(
        "http://user/1",
        Answer{Created: 100, UserId: 1, ExternalId: 500},
        )

    conn.AddAnswer(
        "http://user/2",
        Answer{Created: 50, UserId: 2, ExternalId: 505},
        )
    conn.AddAnswer(
        "http://user/2",
        Answer{Created: 200, UserId: 2, ExternalId: 510},
        )

    answerService = CreateAnswerService(repo, conn)
}

func TestUpdateAndGetUserAnswers(t *testing.T) {
    answerServiceTest()
    userId := 1

    err := answerService.UpdateUserAnswers(userId)

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    answers, _ := answerService.GetUserAnswers(userId)

    if len(answers) != 1 {
        t.Errorf("After updating user answers has %d answers", len(answers))
    }
}

func TestUpdateUserAnswersWillNotCountSameAnswerTwice(t *testing.T) {
    answerServiceTest()
    userId := 2

    err := answerService.UpdateUserAnswers(userId)
    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    err = answerService.UpdateUserAnswers(userId)
    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    answers, _ := answerService.GetUserAnswers(userId)

    if len(answers) != 2 {
        t.Errorf("After updating user answers has %d answers", len(answers))
    }
}

func TestGetUserAnswersSince(t *testing.T) {
    answerServiceTest()
    userId := 2

    err := answerService.UpdateUserAnswers(userId)

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    answers, _ := answerService.GetUserAnswersSince(userId, 205)
    if len(answers) != 0 {
        t.Errorf("Expected 0 answers, got %d answers", len(answers))
    }

    answers, _ = answerService.GetUserAnswersSince(userId, 100)
    if len(answers) != 1 {
        t.Errorf("Expected 1 answer, got %d answers", len(answers))
    }

    answers, _ = answerService.GetUserAnswersSince(userId, 50)
    if len(answers) != 2 {
        t.Errorf("Expected 2 answers, got %d answers", len(answers))
    }
}
