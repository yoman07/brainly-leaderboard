package brainlycrawler

import (
    "testing"
    "fmt"
)

func TestGetUserAnswers(t *testing.T) {
    connector := CreateCrawlerConnector()

    answers, _ := connector.GetUserAnswers("http://zadane.pl/profil/montmorillonit-6680665")
    fmt.Println(answers)
}
