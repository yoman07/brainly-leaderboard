package brainlycrawler

import (
    "golang.org/x/net/html"
    "strings"
    "errors"
    "strconv"
    "net/url"
    "encoding/json"
    "bytes"
)

type ProfileParser struct {
    connector RemoteConnector
}

func CreateProfileParser(c RemoteConnector) *ProfileParser {
    parser := new(ProfileParser)
    parser.connector = c
    return parser
}

func (p *ProfileParser) getUserAnswerDetails(profileUrl string, userId, taskId int) (map[string]string, error) {
    var endpointUrl bytes.Buffer
    siteDomain, _ := p.getSiteFromProfileUrl(profileUrl)
    endpointUrl.WriteString(siteDomain)
    endpointUrl.WriteString("/api/21/api_tasks/main_view/")
    endpointUrl.WriteString(strconv.Itoa(taskId))

    result := make(map[string]string)

    raw, err := p.connector.GetTaskMainViewJson(endpointUrl.String())

    if err != nil {
        return result, err
    }

    decoded := make(map[string]interface{})

    err = json.Unmarshal([]byte(raw), &decoded)

    if err != nil {
        return result, err
    }

    data := decoded["data"].(map[string]interface{})
    responses := data["responses"].([]interface{})

    userResponse, err := getUserReponseDetails(userId, responses)

    if err != nil {
        return result, err
    }

    createdDate := userResponse["created"].(string)
    id := int(userResponse["id"].(float64))

    result["created"] = createdDate
    result["id"] = strconv.Itoa(id)

    return result, nil
}

func getUserReponseDetails(userId int, responses []interface{}) (map[string]interface{}, error) {
    for _, response := range responses {
        details := response.(map[string]interface{})
        detailsUserId := int(details["user_id"].(float64))
        if detailsUserId == userId {
            return details, nil
        }
    }

    return make(map[string]interface{}), errors.New("Response not found for given user")
}

func (p *ProfileParser) getProfileDetails(url string) (map[string]string, error) {
    result := make(map[string]string)

    body, err := p.connector.GetProfileHtml(url)

    if err != nil {
        return result, err
    }

    doc, _ := html.Parse(strings.NewReader(body))

    rankingSpan, _ := getElement(doc, "span", "class", "ranking")
    h2, _ := getTag(rankingSpan, "h2")
    link, _ := getTag(h2, "a")

    result["nick"] = link.FirstChild.Data

    return result, nil
}

func (p *ProfileParser) getAllTasksIdsForUser(url string) ([]int, error) {
    returnIds := make([]int, 0)

    ids, err := p.getTasksIdsFromPage(url)

    if err != nil {
        return returnIds, err
    }

    for _, id := range ids {
        returnIds = append(returnIds, id)
    }

    for i := 2; i < 100; i++ {
        var urlWithPage bytes.Buffer
        urlWithPage.WriteString(url)
        urlWithPage.WriteString("/solved/")
        urlWithPage.WriteString(strconv.Itoa(i))
        ids, _ := p.getTasksIdsFromPage(urlWithPage.String())

        if len(ids) == 0 {
            return returnIds, nil
        }

        for _, id := range ids {
            returnIds = append(returnIds, id)
        }
    }

    return returnIds, nil
}

func (p *ProfileParser) getTasksIdsFromPage(url string) ([]int, error) {
    ids := make([]int, 0)

    body, err := p.connector.GetProfileHtml(url)

    if err != nil {
        return ids, err
    }

    doc, _ := html.Parse(strings.NewReader(body))

    tasksSolved, _ := getElement(doc, "div", "id", "tasks-solved")
    tasksOl, _ := getElement(tasksSolved, "ol", "class", "tasks-list")

    for taskLi := tasksOl.FirstChild; taskLi != nil; taskLi = taskLi.NextSibling {
        taskContent, _ := getElement(taskLi, "div", "class", "task-content")
        link, err := getTag(taskContent, "a")

        if err != nil {
            continue
        }

        linkParts := strings.Split(getAttrValue(link, "href"), "/")

        taskId, _ := strconv.ParseInt(linkParts[2], 10, 0)

        ids = append(ids, int(taskId))
    }

    return ids, nil
}

func (p *ProfileParser) getSiteFromProfileUrl(profileUrl string) (string, error) {
    var site bytes.Buffer

    u, _ := url.Parse(profileUrl)
    site.WriteString(u.Scheme)
    site.WriteString("://")
    site.WriteString(u.Host)

    return site.String(), nil
}

func getElement(n *html.Node, tag, attr, value string) (*html.Node, error) {
    if n.Type == html.ElementNode && n.Data == tag {
        for _, a := range n.Attr {
            if a.Key == attr && a.Val == value {
                return n, nil
                break
            }
        }
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        node, err := getElement(c, tag, attr, value)
        if err == nil {
            return node, nil
        }
    }

    return n, errors.New("Not found")
}

func getTag(n *html.Node, tag string) (*html.Node, error) {
    if n.Type == html.ElementNode && n.Data == tag {
        return n, nil
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        node, err := getTag(c, tag)
        if err == nil {
            return node, nil
        }
    }

    return n, errors.New("Not found")
}

func getAttrValue(n *html.Node, attr string) string {
    for _, a := range n.Attr {
        if a.Key == attr {
            return a.Val
        }
    }
    return ""
}
