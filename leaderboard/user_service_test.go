package leaderboard

import "testing"

var userService *UserService

func userServiceTest() {
    repo := GetRepositoryMock()

    repo.SaveUser(
        User{Id: 1, ProfileUrl: "http://user/1"},
        )

    conn := GetBrainlyConnectorMock()

    userService = CreateUserService(repo, conn)
}

func TestGetExistingUserById(t *testing.T) {
    userServiceTest()
    user, err := userService.GetUser(1)

    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    if user.Id != 1 {
        t.Errorf("Received unexpected user id: %d", user.Id)
    }
}

// func TestGetNotExistingUserById

func TestCreateNewUserSuccess(t *testing.T) {
    userServiceTest()
    name := "John Doe"
    profileUrl := "http://example3.com"
    user, err := userService.CreateUser(name, profileUrl)

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
