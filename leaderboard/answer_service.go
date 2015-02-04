package leaderboard

func CreateAnswerForUser(userId int, created int) (*Answer, error) {
	answer := new(Answer)
	answer.UserId = userId
	answer.Created = created

	return repository.SaveAnswer(answer)
}
