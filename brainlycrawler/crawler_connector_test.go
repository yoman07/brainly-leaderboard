package brainlycrawler

import (
    "testing"
)

var connector *CrawlerConnector

func crawlerConnectorTest() {
    connector = CreateCrawlerConnector(CreateFilesystemRemoteConnector())
}

func TestGetUserAnswers(t *testing.T) {
    crawlerConnectorTest()

    answers, err := connector.GetUserAnswers("http://zadane.pl/profil/montmorillonit-6680665")
    expectedExternalIds := []int{9097505, 9097439, 9097199, 9097025, 9096902,
                      9094857, 9094799, 9094724, 9081419, 9079785}

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    for i, _ := range expectedExternalIds {
        if answers[i].ExternalId != expectedExternalIds[i] {
            t.Errorf("Unexpected value: Expecting: %d, got: %d", expectedExternalIds[i], answers[i].ExternalId)
        }
    }
}

func TestGetUserAnswersForInvalidProfileUrl(t *testing.T) {
    crawlerConnectorTest()

    _, err := connector.GetUserAnswers("invalid-url")

    if err == nil {
        t.Errorf("Unexpected error: %s", err)
    }
}

func TestExtractUserIdFromProfileUrl(t *testing.T) {
    crawlerConnectorTest()

    userId, err := connector.extractUserIdFromProfileUrl("http://brainly.com/profile/NyteRyder-324108")

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    if userId != 324108 {
        t.Errorf("Unexpected userId. Got: %d", userId)
    }
}

func TestConvertTimeToTimestamp(t *testing.T) {
    crawlerConnectorTest()

    timestamp := connector.convertTimeToTimestamp("2015-02-11 13:55:12")

    if timestamp != 1423662912 {
        t.Errorf("Unexpected timestamp. Got: %d", timestamp)
    }
}


