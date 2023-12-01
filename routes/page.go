package routes

import (
	"net/http"
	"text/template"
)

func Page(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	tmpl.Execute(w, nil)
}
