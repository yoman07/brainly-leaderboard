package main

import (
    "github.com/k3nn7/leaderboard/leaderboard"
    "net/http"
    "fmt"
    "encoding/json"
)

func UsersIndex(w http.ResponseWriter, r *http.Request) {
    userIds, _ := application.UserService.GetAllUsersIds()
    users := make([]leaderboard.User, 0)

    for _, userId := range userIds {
        user, _ := application.UserService.GetUser(userId)
        users = append(users, user)
    }

    usersJson, _ := json.Marshal(users)

    fmt.Fprint(w, string(usersJson))
}
