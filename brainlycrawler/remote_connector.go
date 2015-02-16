package brainlycrawler

type RemoteConnector interface {
    GetProfileHtml(url string) (string, error)
    GetTaskMainViewJson(taskId int) (string, error)
}
