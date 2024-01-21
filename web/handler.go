package web

import (
	"net/http"
)

// HomeHandler is the handler for the home route.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := 1
	RenderTemplate(w, r, "home", data)

}

// CreateHandler is the handler for the Create route.
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the creation page."))
}

// PersonHandler is the handler for the Person route.
func PersonHandler(w http.ResponseWriter, r *http.Request) {
	data := 1
	RenderTemplate(w, r, "person", data)
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {

}
