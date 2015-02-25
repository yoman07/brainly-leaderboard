package main

import (
    "github.com/k3nn7/leaderboard/app"
    "github.com/gorilla/mux"
    "net/http"
)

var application app.App

func main() {
    application = app.Wireup()

    go StartWorker()

    router := mux.NewRouter()
    registerHandlers(router)
    http.Handle("/", router)
    http.ListenAndServe(":80", nil)
}
