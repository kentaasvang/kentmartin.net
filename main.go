package main


import (
    "path"
    "net/http"
    "html/template"
    "log"
    "io/ioutil"
)


type Page struct {
	Title string
}


var CurrentDirectory string = "/home/kent/go/src/github.com/kentaasvang/personal_website"


func indexHandler(w http.ResponseWriter, r *http.Request)  {
	page := Page{"Posts"}
	t, err := template.ParseFiles(path.Join(CurrentDirectory, "index.html"))

	if (err != nil) {
		ioutil.WriteFile("go-log-file.log", []byte(err.Error()), 0644)
	}

	t.Execute(w, page)
}


func main() {
    http.HandleFunc("/", indexHandler)
    http.Handle(
	"/vendor/",
	http.StripPrefix("/vendor/", http.FileServer(http.Dir(path.Join(CurrentDirectory, "vendor")))))

    http.Handle(
	"/static/css/",
	http.StripPrefix("/static/css/", http.FileServer(http.Dir(path.Join(CurrentDirectory, "static/css")))))
    http.Handle(
	"/static/js/",
	http.StripPrefix("/static/js/", http.FileServer(http.Dir(path.Join(CurrentDirectory, "static/js")))))

    log.Fatal(http.ListenAndServe(":9990", nil))
}
