package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patrick-rs/htmx_drugseach/routes"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.Page)

	http.ListenAndServe("localhost:6060", r)
}
