package leaderboard

type RepositoryMock struct {
	users   map[int]User
	answers map[int]Answer
}

func (r *RepositoryMock) FindUserById(id int) (*User, error) {
	user := r.users[id]
	return &user, nil
}

func (r *RepositoryMock) SaveUser(user *User) (*User, error) {
	id := len(r.users) + 1
	user.Id = id
	r.users[id] = *user
	return user, nil
}

func (r *RepositoryMock) SaveAnswer(answer *Answer) (*Answer, error) {
	id := len(r.answers) + 1
	answer.Id = id
	r.answers[id] = *answer
	return answer, nil
}

// Initializes RepositoryMock with example user data
func GetRepositoryMock() *RepositoryMock {
	repository := new(RepositoryMock)
	repository.users = make(map[int]User)
	repository.answers = make(map[int]Answer)

	repository.SaveUser(createExampleUser())

	return repository
}

func createExampleUser() *User {
	user := new(User)
	user.ProfileUrl = "http://example.com"
	user.Name = "John Doe"
	user.Nick = "Foo"
	return user
}
