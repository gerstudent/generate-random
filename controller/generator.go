package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gerstudent/generate-random/services"
	"github.com/gorilla/mux"
)

func Generate(w http.ResponseWriter, r *http.Request) {
	id, value := services.GenerateValue()
	json.NewEncoder(w).Encode(map[string]string{
		"id":    id,
		"value": value,
	})
}

func RetrieveValue(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	value, isFound := services.GetValue(id)
	if !isFound {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id not found",
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"id":    id,
		"value": value,
	})
}
