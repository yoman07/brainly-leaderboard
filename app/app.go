package app

import (
    "github.com/k3nn7/leaderboard/leaderboard"
    "github.com/k3nn7/leaderboard/brainlycrawler"
    "time"
)

type App struct {
    UserService *leaderboard.UserService
    AnswerService *leaderboard.AnswerService
    ContestService *leaderboard.ContestService
}

func Wireup() App {
    app := App{}

    repository := leaderboard.GetRepositoryMock()

    repository.SaveUser(leaderboard.User{
        ProfileUrl: "http://zadane.pl/profil/grzegorzostrow-6390443",
        Name: "Grzegorz Ostrowski",
        Nick: "grzegorzostrow",
    })

    repository.SaveUser(leaderboard.User{
        ProfileUrl: "http://zadane.pl/profil/sokpomaranczowy-6387149",
        Name: "Krzysztof Woliński",
        Nick: "sokpomaranczowy",
    })

    repository.SaveUser(leaderboard.User{
        ProfileUrl: "http://zadane.pl/profil/Sprzatacz-6385436",
        Name: "Marta Ryłko",
        Nick: "sprzatacz",
    })

    repository.SaveUser(leaderboard.User{
        ProfileUrl: "http://zadane.pl/profil/montmorillonit-6680665",
        Name: "Aneta Wojdyła",
        Nick: "montorillonit",
    })

    connector := brainlycrawler.CreateCrawlerConnector()

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
