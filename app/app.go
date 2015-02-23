package app

import (
    "github.com/k3nn7/leaderboard/leaderboard"
    "github.com/k3nn7/leaderboard/brainlycrawler"
    "github.com/k3nn7/leaderboard/filesystemrepository"
    "time"
)

type App struct {
    UserService *leaderboard.UserService
    AnswerService *leaderboard.AnswerService
    ContestService *leaderboard.ContestService
}

func Wireup() App {
    app := App{}

    repository, _ := filesystemrepository.CreateRepository("/Users/lukaszlalik/Projects/leaderboard/bin/repository")

    connector := brainlycrawler.CreateCrawlerConnector(brainlycrawler.CreateHttpRemoteConnector())

    app.UserService = leaderboard.CreateUserService(
        repository,
        connector,
    )

    app.AnswerService = leaderboard.CreateAnswerService(
        repository,
        connector,
    )

    app.ContestService = leaderboard.CreateContestService(
        repository,
        app.UserService,
    )

    return app
}

func daysAgoTimestamp(days int) int64 {
    ago := time.Now().AddDate(0, 0, -days)
    return ago.Unix()
}
