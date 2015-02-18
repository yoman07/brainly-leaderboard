package main

import (
    "net/http"
    "fmt"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Ok")
}
