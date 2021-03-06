package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/TrondSpjelkavik/go-API/APIv2/models"
)


func GetLanguages(w http.ResponseWriter, r *http.Request) {

	var languages []models.Languages

	languages = append(languages, models.Languages{ID: 1, Name: "Go"})
	languages = append(languages, models.Languages{ID: 2, Name: "JavaScript"})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(languages)

}