package leaderboard

type BrainlyConnector interface {
    // Gathers user detailed informations
    GetUserDetails(profileUrl string) (map[string]string, error)
    GetUserAnswers(profileUrl string) ([]Answer, error)
}
