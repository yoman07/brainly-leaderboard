package main

import (
    "net/http"
    "fmt"
)

func ScoreIndex(w http.ResponseWriter, r *http.Request) {
    ranking, _ := application.ContestService.GetRanking()


    fmt.Fprintf(w, "Ranking: %v", ranking)
}
