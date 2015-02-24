package filesystemrepository

import (
    "github.com/k3nn7/leaderboard/leaderboard"
    "errors"
    "os"
    "path/filepath"
    "encoding/json"
    "io/ioutil"
)

type Repository struct {
    storageDirectory string
    users   []leaderboard.User
    answers []leaderboard.Answer
    scores []leaderboard.UserScore
    ranking []int
}

func CreateRepository(storageDirectory string) (*Repository, error) {
    repository := new(Repository)

    repository.storageDirectory = storageDirectory
    err := repository.isStorageDirWritable()

    if err != nil {
        return repository, err
    }

    repository.initStorage()

    return repository, nil
}

func (r *Repository) isStorageDirWritable() error {
    info, err := os.Stat(r.storageDirectory)

    if err != nil {
        return err
    }

    if !info.IsDir() {
        return errors.New("Given storage path must be a directory")
    }

    return nil
}

func (r *Repository) initStorage() error {
    r.createUsersRepository()
    r.createAnswersRepository()
    r.createScoresRepository()
    r.createRankingRepository()

    return nil
}

func (r *Repository) createUsersRepository() error {
    repo := "users.json"
    if r.repositoryFileExist(repo) {
        data := r.readRepositortData(repo)
        r.users = make([]leaderboard.User, 0)
        json.Unmarshal(data, &r.users)
        return nil
    }
    r.users = make([]leaderboard.User, 0)
    data, _ := json.Marshal(r.users)
    r.createReposioryFile(repo, data)
    return nil
}

func (r *Repository) createAnswersRepository() error {
    repo := "answers.json"
    if r.repositoryFileExist(repo) {
        data := r.readRepositortData(repo)
        r.answers = make([]leaderboard.Answer, 0)
        json.Unmarshal(data, &r.answers)
        return nil
    }
    r.answers = make([]leaderboard.Answer, 0)
    data, _ := json.Marshal(r.answers)
    r.createReposioryFile(repo, data)
    return nil
}

func (r *Repository) createScoresRepository() error {
    repo := "scores.json"
    if r.repositoryFileExist(repo) {
        data := r.readRepositortData(repo)
        r.scores = make([]leaderboard.UserScore, 0)
        json.Unmarshal(data, &r.scores)
        return nil
    }
    r.scores = make([]leaderboard.UserScore, 0)
    data, _ := json.Marshal(r.scores)
    r.createReposioryFile(repo, data)
    return nil
}

func (r *Repository) createRankingRepository() error {
    repo := "ranking.json"
    if r.repositoryFileExist(repo) {
        data := r.readRepositortData(repo)
        r.ranking = make([]int, 0)
        json.Unmarshal(data, &r.ranking)
        return nil
    }
    r.ranking = make([]int, 0)
    data, _ := json.Marshal(r.ranking)
    r.createReposioryFile(repo, data)
    return nil
}

func (r *Repository) createReposioryFile(file string, data []byte) error {
    f, err := os.Create(filepath.Join(r.storageDirectory, file))
    defer f.Close()
    f.Write(data)
    return err
}

func (r *Repository) repositoryFileExist(file string) bool {
    _, err := os.Stat(filepath.Join(r.storageDirectory, file))
    return err == nil
}

func (r *Repository) readRepositortData(file string) []byte {
    data, _ := ioutil.ReadFile(filepath.Join(r.storageDirectory, file))
    return data
}

func (r *Repository) syncRepositoryToDisk() {
    users, _ := json.Marshal(r.users)
    r.createReposioryFile("users.json", users)

    answers, _ := json.Marshal(r.answers)
    r.createReposioryFile("answers.json", answers)

    scores, _ := json.Marshal(r.scores)
    r.createReposioryFile("scores.json", scores)

    ranking, _ := json.Marshal(r.ranking)
    r.createReposioryFile("ranking.json", ranking)
}

func (r *Repository) FindUserById(id int) (leaderboard.User, error) {
    for _, user := range r.users {
        if user.Id == id {
            return user, nil
        }
    }

    return leaderboard.User{}, errors.New("User with given id not found")
}

func (r *Repository) FindAllUsers() ([]leaderboard.User, error) {
    return r.users, nil
}

func (r *Repository) SaveUser(user leaderboard.User) (leaderboard.User, error) {
    id := len(r.users) + 1
    user.Id = id
    r.users = append(r.users, user)
    r.syncRepositoryToDisk()
    return user, nil
}

func (r *Repository) SaveAnswer(answer leaderboard.Answer) (leaderboard.Answer, error) {
    id := len(r.answers) + 1
    answer.Id = id
    r.answers = append(r.answers, answer)
    r.syncRepositoryToDisk()
    return answer, nil
}

func (r *Repository) FindUserAnswerByExternalId(userId int, externalId int) (leaderboard.Answer, error) {
    for _, answer := range r.answers {
        if answer.UserId == userId && answer.ExternalId == externalId {
            return answer, nil
        }
    }

    return leaderboard.Answer{}, errors.New("User answer with given id not found")
}

func (r *Repository) FindAnswersByUserId(userId int) ([]leaderboard.Answer, error) {
    userAnswers := make([]leaderboard.Answer, 0)
    for _, answer := range r.answers {
        if answer.UserId == userId {
            userAnswers = append(userAnswers, answer)
        }
    }
    return userAnswers, nil
}

func (r *Repository) FindAnswersByUserIdSince(userId int, since int64) ([]leaderboard.Answer, error) {
    userAnswers := make([]leaderboard.Answer, 0)
    for _, answer := range r.answers {
        if answer.UserId == userId && answer.Created >= since {
            userAnswers = append(userAnswers, answer)
        }
    }
    return userAnswers, nil
}

func (r *Repository) SaveUserScore(score leaderboard.UserScore) (leaderboard.UserScore, error) {
    for i, existingScore := range r.scores {
        if existingScore.UserId == score.UserId {
            r.scores[i] = score
            return score, nil
        }
    }

    r.scores = append(r.scores, score)
    r.syncRepositoryToDisk()
    return score, nil
}

func (r *Repository) FindUserScoreByUserId(userId int) (leaderboard.UserScore, error) {
    for _, score := range r.scores {
        if score.UserId == userId {
            return score, nil
        }
    }
    return leaderboard.UserScore{}, errors.New("Given score not found")
}

func (r *Repository) SaveRanking(ranking []int) error {
    r.ranking = ranking
    r.syncRepositoryToDisk()
    return nil
}

func (r *Repository) FindRanking() ([]int, error) {
    return r.ranking, nil
}
