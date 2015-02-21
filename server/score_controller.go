package main

import (
    "net/http"
    "fmt"
    "encoding/json"
)

func ScoreIndex(w http.ResponseWriter, r *http.Request) {
    ranking, _ := application.ContestService.GetRanking()

    rankingJson, _ := json.Marshal(ranking)

    fmt.Fprintf(w, string(rankingJson))
}
