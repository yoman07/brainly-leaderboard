package main

import (
    "github.com/gorilla/mux"
)

func registerHandlers(router *mux.Router) {
    router.HandleFunc("/", Index)
    router.HandleFunc("/users", UsersIndex)
    router.HandleFunc("/users/{userId}/answers", UserAnswers)
    router.HandleFunc("/users/{userId}/score", UserScore)
    router.HandleFunc("/scores", ScoreIndex)
}