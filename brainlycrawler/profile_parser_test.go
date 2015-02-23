package brainlycrawler

import "testing"

var profileParser *ProfileParser

func profileParserTest() {
    profileParser = CreateProfileParser(CreateFilesystemRemoteConnector())
}

func TestGetUserDetails(t *testing.T) {
    profileParserTest()
    profileUrl := "http://zadane.pl/profil/montmorillonit-6680665"

    results, err := profileParser.getProfileDetails(profileUrl)

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    if results["nick"] !=  "montmorillonit" {
        t.Errorf("Nick should be: montmorillonit, got %s", results["nick"])
    }
}

func TestGetUserDetailsInvalidProfileUrl(t *testing.T) {
    profileParserTest()
    profileUrl := "invalid-url"

    _, err := profileParser.getProfileDetails(profileUrl)

    if err == nil {
        t.Errorf("Unexpected error: %s", err)
    }
}

func TestGetTasksIdsFromPage(t *testing.T) {
    profileParserTest()
    profileUrl := "http://zadane.pl/profil/montmorillonit-6680665"

    results, err := profileParser.getTasksIdsFromPage(profileUrl)
    expected := []int{8794339, 8794089, 8794185, 8793666, 8793918}

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    for i, _ := range expected {
        if results[i] != expected[i] {
            t.Errorf("Unexpected value: Expecting: %d, got: %d", expected[i], results[i])
        }
    }
}

func TestGetTasksIdsFromPageInvalidProfileUrl(t *testing.T) {
    profileParserTest()
    profileUrl := "invalid-url"

    _, err := profileParser.getTasksIdsFromPage(profileUrl)

    if err == nil {
        t.Errorf("Unexpected error: %s", err)
    }
}

func TestGetAllTasksIdsForUser(t *testing.T) {
    profileParserTest()
    profileUrl := "http://zadane.pl/profil/montmorillonit-6680665"

    results, err := profileParser.getAllTasksIdsForUser(profileUrl)
    expected := []int{8794339, 8794089, 8794185, 8793666, 8793918,
                      8791911, 8791858, 8786974, 8775172, 8775734,
                  }

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    for i, _ := range expected {
        if results[i] != expected[i] {
            t.Errorf("Unexpected value: Expecting: %d, got: %d", expected[i], results[i])
        }
    }
}

func TestGetAllTasksIdsForUserInvalidProfileUrl(t *testing.T) {
    profileParserTest()
    profileUrl := "invalid-url"

    _, err := profileParser.getAllTasksIdsForUser(profileUrl)

    if err == nil {
        t.Errorf("Unexpected error: %s", err)
    }
}

func TestGetSiteFromProfileUrl(t *testing.T) {
    profileParserTest()
    profileUrl := "http://zadane.pl/profil/montmorillonit-6680665"

    result, err := profileParser.getSiteFromProfileUrl(profileUrl)

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    if result != "http://zadane.pl" {
        t.Errorf("Unexpected result: %s. Should be: http://zadane.pl", result)
    }
}

func TestGetUserAnswerDetailsForTask(t *testing.T) {
    profileParserTest()
    profileUrl := "http://zadane.pl/profil/montmorillonit-6680665"
    userId := 6680665
    taskId := 8791911

    results, err := profileParser.getUserAnswerDetails(profileUrl, userId, taskId)
    created := "2015-02-15 22:35:21"
    externalId := "9094857"

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    if results["created"] !=  created {
        t.Errorf("Expected created date: %s, got: %s", created, results["created"])
    }

    if results["id"] != externalId {
        t.Errorf("Expected id: %s, got: %s", externalId, results["id"])
    }
}

func TestGetUserAnswerDetailsForTaskInvalidProfileUrl(t *testing.T) {
    profileParserTest()
    profileUrl := "invalid-url"
    userId := 6680665
    taskId := 2

    _, err := profileParser.getUserAnswerDetails(profileUrl, userId, taskId)

    if err == nil {
        t.Errorf("Unexpected error: %s", err)
    }
}
