package leaderboard

type BrainlyConnectorMock struct {
}

func (c *BrainlyConnectorMock) GetUserDetails(profileUrl string) (map[string]string, error) {
	details := make(map[string]string)
	details["nick"] = "Foobar"
	return details, nil
}

func GetBrainlyConnectorMock() *BrainlyConnectorMock {
	return new(BrainlyConnectorMock)
}
