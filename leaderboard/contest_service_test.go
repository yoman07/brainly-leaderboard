package leaderboard

import "testing"

var contestService *ContestService

func contestServiceTest() {
    repo := GetRepositoryMock()

    repo.SaveAnswer(Answer{Created: daysAgoTimestamp(3), UserId: 1})
    repo.SaveAnswer(Answer{Created: daysAgoTimestamp(9), UserId: 1})
    repo.SaveAnswer(Answer{Created: daysAgoTimestamp(50), UserId: 1})

    repo.SaveAnswer(Answer{Created: daysAgoTimestamp(1), UserId: 2})
    repo.SaveAnswer(Answer{Created: daysAgoTimestamp(2), UserId: 2})
    repo.SaveAnswer(Answer{Created: daysAgoTimestamp(3), UserId: 2})

    repo.SaveAnswer(Answer{Created: daysAgoTimestamp(59), UserId: 3})
    repo.SaveAnswer(Answer{Created: daysAgoTimestamp(60), UserId: 3})
    repo.SaveAnswer(Answer{Created: daysAgoTimestamp(70), UserId: 3})
    repo.SaveAnswer(Answer{Created: daysAgoTimestamp(80), UserId: 3})

    contestService = CreateContestService(repo)
}

func TestUpdateUserScore(t *testing.T) {
    contestServiceTest()
    userId := 1

    err := contestService.UpdateUserScore(userId)

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    score, _ := contestService.GetUserScore(userId)

    if score.Answers7Days != 1 {
        t.Errorf("Expected 1 answer, got %d answers", score.Answers7Days)
    }

    if score.Answers30Days != 2 {
        t.Errorf("Expected 2 answers, got %d answers", score.Answers30Days)
    }

    if score.Answers90Days != 3 {
        t.Errorf("Expected 3 answers, got %d answers", score.Answers30Days)
    }
}

func TestGetRanking(t *testing.T) {
    contestServiceTest()
    contestService.UpdateUserScore(1)
    contestService.UpdateUserScore(2)
    contestService.UpdateUserScore(3)

    contestService.RecalculateRanking()

    ranking, err := contestService.GetRanking()

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    if ranking[0].UserId != 2 {
        t.Errorf("First place should be for userId: 2, got %s", ranking[0].UserId)
    }

    if ranking[1].UserId != 1 {
        t.Errorf("Second place should be for userId: 1, got %s", ranking[0].UserId)
    }

    if ranking[2].UserId != 3 {
        t.Errorf("Third place should be for userId: 3, got %s", ranking[0].UserId)
    }
}
