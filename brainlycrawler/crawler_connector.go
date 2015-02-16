package brainlycrawler

import "github.com/k3nn7/leaderboard/leaderboard"

type CrawlerConnector struct {

}

func CreateCrawlerConnector() *CrawlerConnector {
    crawler := new(CrawlerConnector)

    return crawler
}

func (c *CrawlerConnector) GetUserAnswers(profileUrl string) ([]leaderboard.Answer, error) {
    // urls, _ getUrlsWithSolvedTasks("http://zadane.pl/profil/montmorillonit-6680665")
    answers := make([]leaderboard.Answer, 0)

    return answers, nil
}
