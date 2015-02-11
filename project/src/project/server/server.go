package server

import (
	"html/template"
	"net/http"
)

type page struct {
	Title string
	Body  string
}

var templates = template.Must(template.ParseFiles("./static/index.html"))

func makeHandler(s string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "index.html", page{Body: s})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Start() {
	http.HandleFunc("/submit", makeHandler("You've submitted a photo!"))
	http.HandleFunc("/", makeHandler("Please submit your photo, and we'll tell you if it is a sunset or not!"))
	http.ListenAndServe(":8080", nil)
}
