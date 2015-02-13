package leaderboard

import (
    "time"
    "sort"
)

type ContestService struct {
    repository Repository
    userService *UserService
}

func CreateContestService(r Repository, u *UserService) *ContestService {
    service := new(ContestService)
    service.repository = r
    service.userService = u

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
    usersIds, _ := c.userService.GetAllUsersIds()
    usersScores := make([]UserScore, 0)

    for _, userId := range usersIds {
        userScore, _ := c.repository.FindUserScoreByUserId(userId)
        usersScores = append(usersScores, userScore)
    }

    sort.Sort(By7DayScore(usersScores))

    sortedUsersIds := make([]int, 0)

    for _, userScore := range usersScores {
        sortedUsersIds = append(sortedUsersIds, userScore.UserId)
    }

    c.repository.SaveRanking(sortedUsersIds)
}

func (c *ContestService) GetUserScore(userId int) (UserScore, error) {
    return c.repository.FindUserScoreByUserId(userId)
}

func (c *ContestService) GetRanking() ([]UserScore, error) {
    sortedUsersIds, _ := c.repository.FindRanking()
    sortedScores := make([]UserScore, 0)

    for _, userId := range sortedUsersIds {
        score, _ := c.GetUserScore(userId)
        sortedScores = append(sortedScores, score)
    }

    return sortedScores, nil
}

func daysAgoTimestamp(days int) int64 {
    ago := time.Now().AddDate(0, 0, -days)
    return ago.Unix()
}

type By7DayScore []UserScore

func (s By7DayScore) Len() int {
    return len(s)
}

func (s By7DayScore) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s By7DayScore) Less(i, j int) bool {
    return s[i].Answers7Days > s[j].Answers7Days
}
