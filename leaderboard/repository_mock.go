package leaderboard

type RepositoryMock struct {
    users   map[int]User
    answers map[int]Answer
}

// Initializes RepositoryMock with example user data
func GetRepositoryMock() *RepositoryMock {
    repository := new(RepositoryMock)
    repository.users = make(map[int]User)
    repository.answers = make(map[int]Answer)

    return repository
}

func (r *RepositoryMock) FindUserById(id int) (User, error) {
    user := r.users[id]
    return user, nil
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

func (r *RepositoryMock) FindAnswersByUserId(userId int) ([]Answer, error) {
    userAnswers := make([]Answer, 0)
    for _, answer := range r.answers {
        if answer.UserId == userId {
            userAnswers = append(userAnswers, answer)
        }
    }
    return userAnswers, nil
}

func (r *RepositoryMock) FindAnswersByUserIdSince(userId int, since int) ([]Answer, error) {
    userAnswers := make([]Answer, 0)
    for _, answer := range r.answers {
        if answer.UserId == userId && answer.Created >= since {
            userAnswers = append(userAnswers, answer)
        }
    }
    return userAnswers, nil
}
