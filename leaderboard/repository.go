package leaderboard

type Repository interface {
    FindUserById(id int) (User, error)
    SaveUser(user User) (User, error)
    FindAllUsers() ([]User, error)

    SaveAnswer(answer Answer) (Answer, error)
    FindUserAnswerByExternalId(userId int, externalId int) (Answer, error)
    FindAnswersByUserId(userId int) ([]Answer, error)
    FindAnswersByUserIdSince(userId int, since int64) ([]Answer, error)

    SaveUserScore(score UserScore) (UserScore, error)
    FindUserScoreByUserId(userId int) (UserScore, error)

    SaveRanking(ranking []int) error
    FindRanking() ([]int, error)
}

