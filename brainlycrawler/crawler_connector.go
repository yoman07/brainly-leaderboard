package brainlycrawler

import (
    "github.com/k3nn7/leaderboard/leaderboard"
    "fmt"
    "strings"
    "strconv"
    "time"
)

type CrawlerConnector struct {
    profileParser *ProfileParser
}

func CreateCrawlerConnector() *CrawlerConnector {
    crawler := new(CrawlerConnector)
    crawler.profileParser = CreateProfileParser(CreateHttpRemoteConnector())
    return crawler
}

func (c *CrawlerConnector) GetUserAnswers(profileUrl string) ([]leaderboard.Answer, error) {
    userId, _ := c.extractUserIdFromProfileUrl(profileUrl)
    urls, _ := c.profileParser.getUrlsWithSolvedTasks(profileUrl)
    answers := make([]leaderboard.Answer, 0)

    taskIds := make([]int, 0)

    for _, url := range urls {
        ids, _ := c.profileParser.getTasksIdsFromPage(url)
        fmt.Println("ids from %s: %v", url, ids)
        for _, taskId := range ids {
            taskIds = append(taskIds, taskId)
        }
    }

    fmt.Println("All ids: ", taskIds)

    for _, taskId := range taskIds {
        var answer leaderboard.Answer
        answerDetails, _ := c.profileParser.getUserAnswerDetails(profileUrl, userId, taskId)
        externalId, _ := strconv.ParseInt(answerDetails["id"], 10, 0)
        answer.ExternalId = int(externalId)
        answer.UserId = userId
        answer.Created = c.convertTimeToTimestamp(answerDetails["created"])
        answers = append(answers, answer)
        fmt.Println(answerDetails)
    }


    return answers, nil
}

func (c *CrawlerConnector) GetUserDetails(profileUrl string) (map[string]string, error) {
    return make(map[string]string), nil
}

func (c *CrawlerConnector) extractUserIdFromProfileUrl(profileUrl string) (int, error) {
    dashPos := strings.LastIndex(profileUrl, "-")
    length := len(profileUrl)
    fmt.Println("Dashpos: ", dashPos, " length: ", length)
    userId := profileUrl[dashPos+1:length]
    intval, _ := strconv.ParseInt(userId, 10, 0)
    return int(intval), nil
}

func (c *CrawlerConnector) convertTimeToTimestamp(timeString string) int64 {
    layout := "2006-01-02 15:04:05"
    timestamp, _ := time.Parse(layout, timeString)
    return timestamp.Unix()
}
