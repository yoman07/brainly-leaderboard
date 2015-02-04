package leaderboard

var repository Repository
var brainlyConnector BrainlyConnector

func SetRepository(r Repository) {
	repository = r
}

func SetBrainlyConnector(c BrainlyConnector) {
	brainlyConnector = c
}

func GetUser(id int) (*User, error) {
	return repository.FindUserById(id)
}

func CreateUser(name string, profileUrl string) (*User, error) {
	user := new(User)
	user.Name = name
	user.ProfileUrl = profileUrl

	userDetails, _ := brainlyConnector.GetUserDetails(profileUrl)
	user.Nick = userDetails["nick"]

	return repository.SaveUser(user)
}
