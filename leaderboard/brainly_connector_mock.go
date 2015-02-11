package leaderboard

type BrainlyConnectorMock struct {
    answers map[string][]Answer
}

func GetBrainlyConnectorMock() *BrainlyConnectorMock {
    connector := new(BrainlyConnectorMock)
    connector.answers = make(map[string][]Answer)

    return connector
}

func (c *BrainlyConnectorMock) GetUserDetails(profileUrl string) (map[string]string, error) {
    details := make(map[string]string)
    details["nick"] = "Foobar"
    return details, nil
}

func (c *BrainlyConnectorMock) GetUserAnswers(profileUrl string) ([]Answer, error) {
    userAnswers := make([]Answer, 0)

    for _, answer := range c.answers[profileUrl] {
        userAnswers = append(userAnswers, answer)
    }

    return userAnswers, nil
}

func (c *BrainlyConnectorMock) AddAnswer(profileUrl string, answer Answer) {
    if _, ok := c.answers[profileUrl]; ok == false {
        c.answers[profileUrl] = make([]Answer, 0)
    }
    c.answers[profileUrl] = append(c.answers[profileUrl], answer)
}
