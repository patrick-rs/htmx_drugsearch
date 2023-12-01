package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patrick-rs/htmx_drugseach/data"
	"github.com/patrick-rs/htmx_drugseach/routes"
)

func main() {

	data.Setup()

	r := mux.NewRouter()

	r.HandleFunc("/", routes.Page)
	r.HandleFunc("/search", routes.Search)

	http.ListenAndServe("localhost:6060", r)
}
