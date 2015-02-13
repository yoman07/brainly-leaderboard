package leaderboard

type UserService struct {
    repository Repository
    connector BrainlyConnector
}

func CreateUserService(r Repository, c BrainlyConnector) *UserService {
    service := new(UserService)
    service.repository = r
    service.connector = c

    return service
}

func (u *UserService) GetUser(id int) (User, error) {
    return u.repository.FindUserById(id)
}

func (u *UserService) CreateUser(name string, profileUrl string) (User, error) {
    user := User{Name: name, ProfileUrl: profileUrl}

    userDetails, _ := u.connector.GetUserDetails(profileUrl)
    user.Nick = userDetails["nick"]

    return u.repository.SaveUser(user)
}

func (u *UserService) GetAllUsersIds() ([]int, error) {
    ids := make([]int, 0)
    users, _ := u.repository.FindAllUsers()

    for _, user := range users {
        ids = append(ids, user.Id)
    }

    return ids, nil
}
