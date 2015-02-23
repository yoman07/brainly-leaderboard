package filesystemrepository

import (
    "testing"
    "os"
    "path/filepath"
)

func TestCreateRepositoryForValidStorageDir(t *testing.T) {
    _, err := CreateRepository("/")

    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
}

func TestCreateRepositoryForInvalidStorageDir(t *testing.T) {
    _, err := CreateRepository("/a/b")

    if err == nil {
        t.Errorf("Unexpected error: %v", err)
    }
}

func TestRepositoryFilesAreCreated(t *testing.T) {
    pwd, _ := os.Getwd()
    repositoryFiles := []string{
        "users.json", "answers.json", "scores.json", "ranking.json",
    }
    repositoryDirectory := filepath.Join(pwd, "repository")

    for _, file := range repositoryFiles {
        removeRepositoryFile(repositoryDirectory, file)
    }

    CreateRepository(repositoryDirectory)

    for _, file := range repositoryFiles {
        if !repositoryFileExist(repositoryDirectory, file) {
            t.Errorf("Repository file does not exists: %v", file)
        }
    }

    for _, file := range repositoryFiles {
        removeRepositoryFile(repositoryDirectory, file)
    }
}

func removeRepositoryFile(dir, file string) {
    os.Remove(filepath.Join(dir, file))
}

func repositoryFileExist(dir, file string) bool {
    _, err := os.Stat(filepath.Join(dir, file))
    return err == nil
}
