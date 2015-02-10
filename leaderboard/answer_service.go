package leaderboard

type AnswerService struct {
	repository       Repository
	brainlyConnector BrainlyConnector
}

func CreateAnswerService(r Repository, b BrainlyConnector) *AnswerService {
	service := new(AnswerService)
	service.repository = r
	service.brainlyConnector = b

	return service
}

func (a *AnswerService) CreateAnswerForUser(userId int, created int) (*Answer, error) {
	answer := new(Answer)
	answer.UserId = userId
	answer.Created = created

	return a.repository.SaveAnswer(answer)
}
