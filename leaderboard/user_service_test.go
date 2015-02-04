package leaderboard

import (
	"testing"
)

func init() {
	SetRepository(GetRepositoryMock())
	SetBrainlyConnector(GetBrainlyConnectorMock())
}

func TestGetExistingUserById(t *testing.T) {
	user, err := GetUser(1)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if user.Id != 1 {
		t.Errorf("Received unexpected user id: %d", user.Id)
	}
}

// func TestGetNotExistingUserById

func TestCreateNewUserSuccess(t *testing.T) {
	name := "John Doe"
	profileUrl := "http://example3.com"
	user, err := CreateUser(name, profileUrl)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if user.Id != 2 {
		t.Errorf("Unexpected id: %d", user.Id)
	}

	if user.ProfileUrl != profileUrl {
		t.Errorf("Unexpected profileUrl: %s", user.ProfileUrl)
	}

	if user.Name != name {
		t.Errorf("Unexpected name: %s", user.Name)
	}

	if user.Nick != "Foobar" {
		t.Errorf("Unexpected nick: %s", user.Nick)
	}
}
