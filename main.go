package main

import (
	"os"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	_ "github.com/joho/godotenv/autoload"
)


type Page struct {
	Title string
}

// load project root 
var Root string = os.Getenv("PROJECT_PATH")


func indexHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Kentmartin | Hjem"}

	t, err := template.ParseFiles(path.Join(Root, "index.html"))

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
			http.Dir(path.Join(Root, "vendor")))))

	// Custom CSS
	http.Handle(
		"/static/css/",
		http.StripPrefix("/static/css/", http.FileServer(
			http.Dir(path.Join(Root, "static/css")))))

	// Custom JS
	http.Handle(
		"/static/js/",
		http.StripPrefix("/static/js/", http.FileServer(
			http.Dir(path.Join(Root, "static/js")))))

	log.Fatal(http.ListenAndServe(":9990", nil))
}
