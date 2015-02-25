package main

import (
    "github.com/gorilla/mux"
)

func registerHandlers(router *mux.Router) {
    router.HandleFunc("/", Index)
    router.HandleFunc("/js/{fileName}", Js)
    router.HandleFunc("/css/{fileName}", Css)
    router.HandleFunc("/img/{fileName}", Img)
    router.HandleFunc("/users", UsersIndex)
    router.HandleFunc("/users/{userId}/answers", UserAnswers)
    router.HandleFunc("/users/{userId}/score", UserScore)
    router.HandleFunc("/scores", ScoreIndex)
}
