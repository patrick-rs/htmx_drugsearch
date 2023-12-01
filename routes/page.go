package routes

import (
	"net/http"
	"text/template"

	"github.com/patrick-rs/htmx_drugseach/data"
)

func Page(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	tmpl.Execute(w, nil)
}

func Search(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("search")

	matchingDrugs := data.SearchDrugs(s)

	tmpl := template.Must(template.ParseFiles("index.html"))

	err := tmpl.ExecuteTemplate(w, "Drugs", matchingDrugs)

	if err != nil {
		panic(err)
	}
}
