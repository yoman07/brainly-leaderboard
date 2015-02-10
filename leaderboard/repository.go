package leaderboard

type Repository interface {
	FindUserById(id int) (*User, error)
	SaveUser(user *User) (*User, error)

	SaveAnswer(answer *Answer) (*Answer, error)
}