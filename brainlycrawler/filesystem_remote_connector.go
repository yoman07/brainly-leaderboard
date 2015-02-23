package brainlycrawler

import (
    "bytes"
    "os"
    "io/ioutil"
    "strings"
    "strconv"
    "errors"
)

type FilesystemRemoteConnector struct {

}

func CreateFilesystemRemoteConnector() *FilesystemRemoteConnector {
    return new(FilesystemRemoteConnector)
}

func (c *FilesystemRemoteConnector) GetProfileHtml(url string) (string, error) {
    var profilePath bytes.Buffer
    pwd, _ := os.Getwd()

    profilePath.WriteString(pwd)

    if url == "http://zadane.pl/profil/montmorillonit-6680665" {
        profilePath.WriteString("/assets/montmorillonit.html")
    } else if url == "http://zadane.pl/profil/montmorillonit-6680665/solved/2" {
        profilePath.WriteString("/assets/montmorillonit-2.html")
    } else {
        return "", errors.New("Invalid URL")
    }


    html, _ := ioutil.ReadFile(profilePath.String())

    return string(html), nil
}

func (c *FilesystemRemoteConnector) GetTaskMainViewJson(url string) (string, error) {
    var endpointPath bytes.Buffer
    pwd, _ := os.Getwd()

    taskId, _ := c.extractTaskIdFromUrl(url)

    endpointPath.WriteString(pwd)
    endpointPath.WriteString("/assets/")
    endpointPath.WriteString(strconv.Itoa(taskId))
    endpointPath.WriteString(".json")

    json, err := ioutil.ReadFile(endpointPath.String())

    if err != nil {
        return "", err
    }

    return string(json), nil
}

func (c *FilesystemRemoteConnector) extractTaskIdFromUrl(url string) (int, error) {
    dashPos := strings.LastIndex(url, "/")
    length := len(url)
    taskId := url[dashPos+1:length]
    intval, _ := strconv.ParseInt(taskId, 10, 0)
    return int(intval), nil
}
