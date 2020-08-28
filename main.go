package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

type Page struct {
	Title string
}

// TODO: this should _NOT_ be hardcoded
// TODO: this should _NOT_ be hardcoded
// TODO: this should _NOT_ be hardcoded
var CurrentDirectory string =
	"/Users/kent/go/src/github.com/kentaasvang/personal_website"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Posts"}
	t, err := template.ParseFiles(path.Join(CurrentDirectory, "index.html"))

	if err != nil {
		ioutil.WriteFile("go-log-file.log", []byte(err.Error()), 0644)
	}

	t.Execute(w, page)
}

func main() {

	// Routes
	http.HandleFunc("/", indexHandler)

	// jQuery, bootstrap etc
	http.Handle(
		"/vendor/",
		http.StripPrefix("/vendor/", http.FileServer(
			http.Dir(path.Join(CurrentDirectory, "vendor")))))

	// Custom CSS
	http.Handle(
		"/static/css/",
		http.StripPrefix("/static/css/", http.FileServer(
			http.Dir(path.Join(CurrentDirectory, "static/css")))))

	// Custom JS
	http.Handle(
		"/static/js/",
		http.StripPrefix("/static/js/", http.FileServer(
			http.Dir(path.Join(CurrentDirectory, "static/js")))))

	log.Fatal(http.ListenAndServe(":9990", nil))
}
