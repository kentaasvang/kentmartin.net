package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Shit is working")
    })

    http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to the admin")
    })

    log.Fatal(http.ListenAndServe(":9990", nil))

}
