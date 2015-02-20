package brainlycrawler

import (
    "github.com/k3nn7/leaderboard/leaderboard"
    "strings"
    "strconv"
    "time"
)

type CrawlerConnector struct {
    profileParser *ProfileParser
}

func CreateCrawlerConnector(r RemoteConnector) *CrawlerConnector {
    crawler := new(CrawlerConnector)
    crawler.profileParser = CreateProfileParser(r)
    return crawler
}

func (c *CrawlerConnector) GetUserAnswers(profileUrl string) ([]leaderboard.Answer, error) {
    userId, _ := c.extractUserIdFromProfileUrl(profileUrl)
    answers := make([]leaderboard.Answer, 0)

    taskIds := make([]int, 0)

    taskIds, _ = c.profileParser.getAllTasksIdsForUser(profileUrl)

    for _, taskId := range taskIds {
        var answer leaderboard.Answer
        answerDetails, _ := c.profileParser.getUserAnswerDetails(profileUrl, userId, taskId)
        externalId, _ := strconv.ParseInt(answerDetails["id"], 10, 0)
        answer.ExternalId = int(externalId)
        answer.UserId = userId
        answer.Created = c.convertTimeToTimestamp(answerDetails["created"])
        answers = append(answers, answer)
    }


    return answers, nil
}

func (c *CrawlerConnector) GetUserDetails(profileUrl string) (map[string]string, error) {
    return make(map[string]string), nil
}

func (c *CrawlerConnector) extractUserIdFromProfileUrl(profileUrl string) (int, error) {
    dashPos := strings.LastIndex(profileUrl, "-")
    length := len(profileUrl)
    userId := profileUrl[dashPos+1:length]
    intval, _ := strconv.ParseInt(userId, 10, 0)
    return int(intval), nil
}

func (c *CrawlerConnector) convertTimeToTimestamp(timeString string) int64 {
    layout := "2006-01-02 15:04:05"
    timestamp, _ := time.Parse(layout, timeString)
    return timestamp.Unix()
}
