package main

import (
	"net/http"

	"github.com/gerstudent/generate-random/controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/generate/", controller.Generate).Methods("POST")
	r.HandleFunc("/api/retrieve/{id}", controller.RetrieveValue).Methods("GET")

	http.ListenAndServe(":8081", r)
}
