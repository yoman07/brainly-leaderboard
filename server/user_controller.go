package main

import (
    "github.com/k3nn7/leaderboard/leaderboard"
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
    "strconv"
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

    fmt.Fprintf(w, string(usersJson))
}

func UserAnswers(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId64, _ := strconv.ParseInt(vars["userId"], 10, 0)
    userId := int(userId64)

    answers, _ := application.AnswerService.GetUserAnswers(userId)

    fmt.Fprintf(w, "User: %d answers: %v", userId, answers)
}

func UserScore(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userId64, _ := strconv.ParseInt(vars["userId"], 10, 0)
    userId := int(userId64)

    score, _ := application.ContestService.GetUserScore(userId)

    fmt.Fprintf(w, "User: %d score: %v", userId, score)
}
