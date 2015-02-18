package main

import (
    "github.com/k3nn7/leaderboard/leaderboard"
    "net/http"
    "fmt"
)

func AnswersView(w http.ResponseWriter, r *http.Request) {
    userIds, _ := application.UserService.GetAllUsersIds()
    users := make([]leaderboard.User, 0)

    for _, userId := range userIds {
        user, _ := application.UserService.GetUser(userId)
        users = append(users, user)
    }

    fmt.Fprintf(w, "Users: %v", users)
}
