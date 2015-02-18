package main

import (
    "time"
    "fmt"
)

func StartWorker() {
    tick := time.Tick(5 * time.Second)
    op := 0

    for {
        select {
        case <- tick:
            op = runOp(op)
        default:
            time.Sleep(100 * time.Millisecond)
        }
    }
}

func runOp(op int) int {
    switch op {
    case 0:
        updateUsersAnswers()
        return 1
    case 1:
        updateUsersScore()
        return 2
    case 2:
        recalculateRanking()
        return 0
    }

    return 0
}

func updateUsersAnswers() {
    usersIds, _ := application.UserService.GetAllUsersIds()
    for _, userId := range usersIds {
        fmt.Println("Update answer data for user: %d", userId)
        application.AnswerService.UpdateUserAnswers(userId)
    }
}

func updateUsersScore() {
    usersIds, _ := application.UserService.GetAllUsersIds()
    for _, userId := range usersIds {
        fmt.Println("Update score for user: %d", userId)
        application.ContestService.UpdateUserScore(userId)
    }
}

func recalculateRanking() {
    fmt.Println("recalculateRanking")
    application.ContestService.RecalculateRanking()
}
