package brainlycrawler

import (
    "net/http"
    "io/ioutil"
)

type HttpRemoteConnector struct {

}

func CreateHttpRemoteConnector() *HttpRemoteConnector {
    return new(HttpRemoteConnector)
}

func (c *HttpRemoteConnector) GetProfileHtml(url string) (string, error) {
    resp, _ := http.Get(url)
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    return string(body), nil
}

func (c *HttpRemoteConnector) GetTaskMainViewJson(url string) (string, error) {
    resp, _ := http.Get(url)
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    return string(body), nil
}
