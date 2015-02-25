package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
    "os"
    "io/ioutil"
    "bytes"
)

func Index(w http.ResponseWriter, r *http.Request) {
    var filePath bytes.Buffer
    pwd, _ := os.Getwd()

    filePath.WriteString(pwd)
    filePath.WriteString("/../assets/index.html")

    content, _ := ioutil.ReadFile(filePath.String())

    fmt.Fprint(w, string(content))
}

func Js(w http.ResponseWriter, r *http.Request) {
    var filePath bytes.Buffer
    pwd, _ := os.Getwd()
    vars := mux.Vars(r)

    filePath.WriteString(pwd)
    filePath.WriteString("/../assets/js/")
    filePath.WriteString(vars["fileName"])

    content, _ := ioutil.ReadFile(filePath.String())

    w.Header().Set("Content-type", "text/javascript")
    fmt.Fprint(w, string(content))
}

func Css(w http.ResponseWriter, r *http.Request) {
    var filePath bytes.Buffer
    pwd, _ := os.Getwd()
    vars := mux.Vars(r)

    filePath.WriteString(pwd)
    filePath.WriteString("/../assets/css/")
    filePath.WriteString(vars["fileName"])

    content, _ := ioutil.ReadFile(filePath.String())

    w.Header().Set("Content-type", "text/css")
    fmt.Fprint(w, string(content))
}

func Img(w http.ResponseWriter, r *http.Request) {
    var filePath bytes.Buffer
    pwd, _ := os.Getwd()
    vars := mux.Vars(r)

    filePath.WriteString(pwd)
    filePath.WriteString("/../assets/img/")
    filePath.WriteString(vars["fileName"])

    content, _ := ioutil.ReadFile(filePath.String())

    w.Header().Set("Content-type", "image/png")
    fmt.Fprint(w, string(content))
}
