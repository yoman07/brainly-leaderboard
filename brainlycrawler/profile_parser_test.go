package brainlycrawler

import "testing"

var profileParser *ProfileParser

func profileParserTest() {
    profileParser = CreateProfileParser()
}

func TestGetUserDetails(t *testing.T) {
    profileUrl := "http://zadane.pl/profil/montmorillonit-6680665"

    results, err := profileParser.getProfileDetails(profileUrl)

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    if results["nick"] !=  "montmorillonit" {
        t.Errorf("Nick should be: montmorillonit, got %s", results["nick"])
    }
}

func TestGetTasksIdsFromPage(t *testing.T) {
    profileUrl := "http://zadane.pl/profil/montmorillonit-6680665"

    results, err := profileParser.getTasksIdsFromPage(profileUrl)
    expected := []int{8775172, 8775734, 8774120, 8774146, 8769141}

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    for i, _ := range expected {
        if results[i] != expected[i] {
            t.Errorf("Unexpected value: Expecting: %d, got: %d", expected[i], results[i])
        }
    }
}

func TestGetSiteFromProfileUrl(t *testing.T) {
    profileUrl := "http://zadane.pl/profil/montmorillonit-6680665"

    result, err := profileParser.getSiteFromProfileUrl(profileUrl)

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    if result != "http://zadane.pl" {
        t.Errorf("Unexpected result: %s. Should be: http://zadane.pl", result)
    }
}

func TestGetTasksSolvedPagesUrls(t *testing.T) {
    profileUrl := "http://zadane.pl/profil/montmorillonit-6680665"

    results, err := profileParser.getUrlsWithSolvedTasks(profileUrl)
    expected := []string{
        "http://zadane.pl/profil/montmorillonit-6680665",
        "http://zadane.pl/profil/montmorillonit-6680665/solved/2",
        "http://zadane.pl/profil/montmorillonit-6680665/solved/3",
        "http://zadane.pl/profil/montmorillonit-6680665/solved/4",
    }

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    for i, _ := range expected {
        if results[i] != expected[i] {
            t.Errorf("Unexpected value: Expecting: %s, got: %s", expected[i], results[i])
        }
    }
}

func TestGetUserAnswerDetailsForTask(t *testing.T) {
    userId := 6680665
    taskId := 8791911

    results, err := profileParser.getUserAnswerDetails(userId, taskId)
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
