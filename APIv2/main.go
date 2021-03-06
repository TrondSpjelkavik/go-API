package main

import (
	"log"
	"net/http"

	"github.com/TrondSpjelkavik/go-API/APIv2/handlers"
	"github.com/go-chi/chi/v5"
)


func main() {

	router := chi.NewRouter()
	router.Get("/api/languages", handlers.GetLanguages)

	err := http.ListenAndServe(":2102", router)
	if err != nil {
		log.Fatal(err)
	}

}
