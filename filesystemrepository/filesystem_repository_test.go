package filesystemrepository

import (
    "testing"
    "os"
    "path/filepath"
    "github.com/k3nn7/leaderboard/leaderboard"
)

var repository *Repository

func setupRepositoryTest() {
    repositoryFiles := []string{
        "users.json", "answers.json", "scores.json", "ranking.json",
    }
    pwd, _ := os.Getwd()
    repositoryDirectory := filepath.Join(pwd, "repository")

    for _, file := range repositoryFiles {
        removeRepositoryFile(repositoryDirectory, file)
    }
    repository, _ = CreateRepository(repositoryDirectory)
}

func cleanupRepositoryTest() {
    repositoryFiles := []string{
        "users.json", "answers.json", "scores.json", "ranking.json",
    }
    pwd, _ := os.Getwd()
    repositoryDirectory := filepath.Join(pwd, "repository")

    for _, file := range repositoryFiles {
        removeRepositoryFile(repositoryDirectory, file)
    }
}

func TestSaveUser(t *testing.T) {
    setupRepositoryTest()

    user := leaderboard.User{ProfileUrl: "http://example.com"}

    createdUser, err := repository.SaveUser(user)

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if createdUser.ProfileUrl != user.ProfileUrl {
        t.Errorf(
            "Invalid user ProfileUrl. Expecting: %v, got: %v",
            user.ProfileUrl,
            createdUser.ProfileUrl,
        )
    }

    cleanupRepositoryTest()
}

func TestFindUserById(t *testing.T) {
    setupRepositoryTest()
    user1url := "http://example.com"
    user2url := "http://example2.com"

    repository.SaveUser(leaderboard.User{ProfileUrl: user1url})
    repository.SaveUser(leaderboard.User{ProfileUrl: user2url})

    user, err := repository.FindUserById(1)

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if user.ProfileUrl != user1url {
        t.Errorf(
            "Invalid user ProfileUrl. Expecting: %v, got: %v",
            user1url,
            user.ProfileUrl,
        )
    }

    user, err = repository.FindUserById(2)

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if user.ProfileUrl != user2url {
        t.Errorf(
            "Invalid user ProfileUrl. Expecting: %v, got: %v",
            user2url,
            user.ProfileUrl,
        )
    }

    cleanupRepositoryTest()
}
