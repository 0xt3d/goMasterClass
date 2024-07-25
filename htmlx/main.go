package main

import (
    "net/http"
    "github.com/mjibson/esc"
)

//go:generate esc -o templates.go -pkg main templates.htmlx

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := struct {
            Title   string
            Message string
        }{
            Title:   "Simple Go HTMLX Example",
            Message: "Hello, World!",
        }
        if err := templates.Index.Execute(w, data); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

    http.ListenAndServe(":8080", nil)
}
