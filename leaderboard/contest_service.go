package leaderboard

import (
    "time"
    )

type ContestService struct {
    repository Repository
}

func CreateContestService(r Repository) *ContestService {
    service := new(ContestService)
    service.repository = r

    return service
}

func (c *ContestService) UpdateUserScore(userId int) error {
    answers7Days, _ := c.repository.FindAnswersByUserIdSince(
        userId,
        daysAgoTimestamp(7),
    )

    answers30Days, _ := c.repository.FindAnswersByUserIdSince(
        userId,
        daysAgoTimestamp(30),
    )

    answers90Days, _ := c.repository.FindAnswersByUserIdSince(
        userId,
        daysAgoTimestamp(90),
    )

    score := UserScore{
        UserId: userId,
        Answers7Days: len(answers7Days),
        Answers30Days: len(answers30Days),
        Answers90Days: len(answers90Days),
    }

    c.repository.SaveUserScore(score)

    return nil
}

func (c *ContestService) RecalculateRanking() {
}

func (c *ContestService) GetUserScore(userId int) (UserScore, error) {
    return c.repository.FindUserScoreByUserId(userId)
}

func daysAgoTimestamp(days int) int64 {
    ago := time.Now().AddDate(0, 0, -days)
    return ago.Unix()
}
