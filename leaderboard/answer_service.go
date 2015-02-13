package leaderboard

type AnswerService struct {
    repository       Repository
    connector BrainlyConnector
}

func CreateAnswerService(r Repository, b BrainlyConnector) *AnswerService {
    service := new(AnswerService)
    service.repository = r
    service.connector = b

    return service
}

func (a *AnswerService) UpdateUserAnswers(userId int) error {
    user, _ := a.repository.FindUserById(userId)
    answers, _ := a.connector.GetUserAnswers(user.ProfileUrl)

    for _, answer := range answers {
        a.repository.SaveAnswer(answer)
    }

    return nil
}

func (a *AnswerService) GetUserAnswers(userId int) ([]Answer, error) {
    answers, _ := a.repository.FindAnswersByUserId(userId)
    return answers, nil
}

func (a *AnswerService) GetUserAnswersSince(userId int, since int64) ([]Answer, error) {
    answers, _ := a.repository.FindAnswersByUserIdSince(userId, since)
    return answers, nil
}
