package brainlycrawler

import (
    "io/ioutil"
    "os"
    "bytes"
    "golang.org/x/net/html"
    "strings"
    "errors"
    "strconv"
    "net/url"
    "encoding/json"
)

func getUserAnswerDetails(userId, taskId int) (map[string]string, error) {
    result := make(map[string]string)

    raw, _ := getTaskMainViewJson(taskId)

    decoded := make(map[string]interface{})

    json.Unmarshal([]byte(raw), &decoded)
    data := decoded["data"].(map[string]interface{})
    responses := data["responses"].([]interface{})

    userResponse, _ := getUserReponseDetails(userId, responses)
    createdDate := userResponse["created"].(string)

    result["created"] = createdDate


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

    return make(map[string]interface{}), nil
}

func getProfileDetails(url string) (map[string]string, error) {
    result := make(map[string]string)

    body, _ := getProfileHtml(url)
    doc, _ := html.Parse(strings.NewReader(body))

    rankingSpan, _ := getElement(doc, "span", "class", "ranking")
    h2, _ := getTag(rankingSpan, "h2")
    link, _ := getTag(h2, "a")

    result["nick"] = link.FirstChild.Data

    return result, nil
}

func getTasksIdsFromPage(url string) ([]int, error) {
    ids := make([]int, 0)

    body, _ := getProfileHtml(url)
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

func getUrlsWithSolvedTasks(profileUrl string) ([]string, error) {
    urls := make([]string, 0)
    urls = append(urls, profileUrl)

    body, _ := getProfileHtml(profileUrl)
    doc, _ := html.Parse(strings.NewReader(body))

    tasksSolved, _ := getElement(doc, "div", "id", "tasks-solved")
    pager, _ := getElement(tasksSolved, "div", "class", "pager")

    siteUrl, _ := getSiteFromProfileUrl(profileUrl)

    for span := pager.FirstChild; span != nil; span = span.NextSibling {
        if getAttrValue(span, "class") == "current" {
            continue
        }

        if span.Data != "span" {
            continue
        }

        link, _ := getTag(span, "a")

        var fullUrl bytes.Buffer
        pageUrl := getAttrValue(link, "href")
        fullUrl.WriteString(siteUrl)
        fullUrl.WriteString(pageUrl)

        urls = append(urls, fullUrl.String())
    }

    return urls, nil
}

func getSiteFromProfileUrl(profileUrl string) (string, error) {
    var site bytes.Buffer

    u, _ := url.Parse(profileUrl)
    site.WriteString(u.Scheme)
    site.WriteString("://")
    site.WriteString(u.Host)

    return site.String(), nil
}

func getProfileHtml(url string) (string, error) {
    var profilePath bytes.Buffer
    pwd, _ := os.Getwd()

    profilePath.WriteString(pwd)
    profilePath.WriteString("/assets/montmorillonit.html")

    html, _ := ioutil.ReadFile(profilePath.String())

    return string(html), nil
}

func getTaskMainViewJson(taskId int) (string, error) {
    var endpointPath bytes.Buffer
    pwd, _ := os.Getwd()

    endpointPath.WriteString(pwd)
    endpointPath.WriteString("/assets/tasks.main_view.json")

    json, _ := ioutil.ReadFile(endpointPath.String())

    return string(json), nil
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
