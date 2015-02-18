package brainlycrawler

import (
    "testing"
)

func TestGetUserAnswers(t *testing.T) {
    // connector := CreateCrawlerConnector()

    // answers, _ := connector.GetUserAnswers("http://brainly.com/profile/NyteRyder-324108", 324108)
    // fmt.Println(answers)
}

func TestExtractUserIdFromProfileUrl(t *testing.T) {
    connector := CreateCrawlerConnector()

    userId, err := connector.extractUserIdFromProfileUrl("http://brainly.com/profile/NyteRyder-324108")

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    if userId != 324108 {
        t.Errorf("Unexpected userId. Got: %d", userId)
    }
}

func TestConvertTimeToTimestamp(t *testing.T) {
    connector := CreateCrawlerConnector()

    timestamp := connector.convertTimeToTimestamp("2015-02-11 13:55:12")

    if timestamp != 1423662912 {
        t.Errorf("Unexpected timestamp. Got: %d", timestamp)
    }
}


