package main


import (
    "net/http"
    "log"
    "io/ioutil"
    "html/template"
)

type Page struct {
	Title string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Posts"}
	template, err := template.ParseFiles("~/go/src/github.com/kentaasvang/personal_website/index.html")
	if (err != nil) {
		ioutil.WriteFile("go-log.log", []byte(err.Error()), 0644)
	}
	template.Execute(w, page)
}

func main() {
    http.HandleFunc("/", indexHandler)
    log.Fatal(http.ListenAndServe(":9990", nil))
}
