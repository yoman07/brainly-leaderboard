package leaderboard

import "errors"

type RepositoryMock struct {
    users   map[int]User
    answers map[int]Answer
    scores map[int]UserScore
    ranking []int
}

// Initializes RepositoryMock with example user data
func GetRepositoryMock() *RepositoryMock {
    repository := new(RepositoryMock)
    repository.users = make(map[int]User)
    repository.answers = make(map[int]Answer)
    repository.scores = make(map[int]UserScore)
    repository.ranking = make([]int, 0)

    return repository
}

func (r *RepositoryMock) FindUserById(id int) (User, error) {
    user := r.users[id]
    return user, nil
}

func (r *RepositoryMock) FindAllUsers() ([]User, error) {
    users := make([]User, 0)
    for _, user := range r.users {
        users = append(users, user)
    }
    return users, nil
}

func (r *RepositoryMock) SaveUser(user User) (User, error) {
    id := len(r.users) + 1
    user.Id = id
    r.users[id] = user
    return user, nil
}

func (r *RepositoryMock) SaveAnswer(answer Answer) (Answer, error) {
    id := len(r.answers) + 1
    answer.Id = id
    r.answers[id] = answer

    return answer, nil
}

func (r *RepositoryMock) FindUserAnswerByExternalId(userId int, externalId int) (Answer, error) {
    for _, answer := range r.answers {
        if answer.UserId == userId && answer.ExternalId == externalId {
            return answer, nil
        }
    }

    return Answer{}, errors.New("User answer with given id not found")
}

func (r *RepositoryMock) FindAnswersByUserId(userId int) ([]Answer, error) {
    userAnswers := make([]Answer, 0)
    for _, answer := range r.answers {
        if answer.UserId == userId {
            userAnswers = append(userAnswers, answer)
        }
    }
    return userAnswers, nil
}

func (r *RepositoryMock) FindAnswersByUserIdSince(userId int, since int64) ([]Answer, error) {
    userAnswers := make([]Answer, 0)
    for _, answer := range r.answers {
        if answer.UserId == userId && answer.Created >= since {
            userAnswers = append(userAnswers, answer)
        }
    }
    return userAnswers, nil
}

func (r *RepositoryMock) SaveUserScore(score UserScore) (UserScore, error) {
    userId := score.UserId
    r.scores[userId] = score
    return score, nil
}

func (r *RepositoryMock) FindUserScoreByUserId(userId int) (UserScore, error) {
    score := r.scores[userId]
    return score, nil
}

func (r *RepositoryMock) SaveRanking(ranking []int) error {
    r.ranking = ranking
    return nil
}

func (r *RepositoryMock) FindRanking() ([]int, error) {
    return r.ranking, nil
}
