package main

import (
	"os"
	"github.com/joho/godotenv"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)


type Page struct {
	Title string
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Kentmartin | Hjem"}
	t, err := template.ParseFiles(path.Join(os.Getenv("PROJECT_PATH"), "index.html"))

	if err != nil {
		ioutil.WriteFile("go-log-file.log", []byte(err.Error()), 0644)
	}

	t.Execute(w, page)
}

func main() {
	err := godotenv.Load()
	if err != nil { log.Fatal("Error loading .env file") }

	// Routes
	http.HandleFunc("/", indexHandler)

	// jQuery, bootstrap etc
	http.Handle(
		"/vendor/",
		http.StripPrefix("/vendor/", http.FileServer(
			http.Dir(path.Join(os.Getenv("PROJECT_PATH"), "vendor")))))

	// Custom CSS
	http.Handle(
		"/static/css/",
		http.StripPrefix("/static/css/", http.FileServer(
			http.Dir(path.Join(os.Getenv("PROJECT_PATH"), "static/css")))))

	// Custom JS
	http.Handle(
		"/static/js/",
		http.StripPrefix("/static/js/", http.FileServer(
			http.Dir(path.Join(os.Getenv("PROJECT_PATH"), "static/js")))))

	log.Fatal(http.ListenAndServe(":9990", nil))
}
