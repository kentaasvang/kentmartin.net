package main


import (
    "net/http"
    "html/template"
    "log"
)

type Page struct {
	Title string
}


func indexHandler(w http.ResponseWriter, r *http.Request)  {
	page := Page{"Posts"}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, page)
}


func main() {
    http.HandleFunc("/", indexHandler)
    log.Fatal(http.ListenAndServe(":9990", nil))
}
