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
    resp, err := http.Get(url)

    if err != nil {
        return "", err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        return "", err
    }

    return string(body), nil
}

func (c *HttpRemoteConnector) GetTaskMainViewJson(url string) (string, error) {
    resp, err := http.Get(url)

    if err != nil {
        return "", err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        return "", err
    }

    return string(body), nil
}
