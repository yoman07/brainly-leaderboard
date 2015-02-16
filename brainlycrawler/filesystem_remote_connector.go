package brainlycrawler

import (
    "bytes"
    "os"
    "io/ioutil"
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
    profilePath.WriteString("/assets/montmorillonit.html")

    html, _ := ioutil.ReadFile(profilePath.String())

    return string(html), nil
}

func (c *FilesystemRemoteConnector) GetTaskMainViewJson(taskId int) (string, error) {
    var endpointPath bytes.Buffer
    pwd, _ := os.Getwd()

    endpointPath.WriteString(pwd)
    endpointPath.WriteString("/assets/tasks.main_view.json")

    json, _ := ioutil.ReadFile(endpointPath.String())

    return string(json), nil
}
