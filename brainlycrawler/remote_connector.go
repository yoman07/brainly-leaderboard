package brainlycrawler

type RemoteConnector interface {
    GetProfileHtml(url string) (string, error)
    GetTaskMainViewJson(url string) (string, error)
}
