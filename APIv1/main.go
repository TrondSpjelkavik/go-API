package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)



func get(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"GET": "Gets Trond Fuglseth Spjelkavik"}`))

}

func post(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"POST": "Posts Trond Fuglseth Spjelkavik"}`))

}


func put(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"PUT": "Updates Trond Fuglseth Spjelkavik"}`))

}

func delete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"DELETE": "Deletes Trond Fuglseth Spjelkavik"}`))

}

func notFound(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"NOT FOUND": "Did not find Trond Fuglseth Spjelkavik"}`))

}

func params(w http.ResponseWriter, r *http.Request) {

	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	userID := -1
	var err error
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"Error": "need a number"}`))
			return
		}
	}

	commentID := -1
	if val, ok := pathParams["comentID"]; ok {
		commentID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"Error": "need a number"}`))
			return
		}
	}

	query := r.URL.Query()
	location := query.Get("location")

	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))

}


func main() {

	route := mux.NewRouter()

	api := route.PathPrefix("/trondApi/version1").Subrouter()

	api.HandleFunc("/", get).Methods(http.MethodGet)
	api.HandleFunc("/", post).Methods(http.MethodPost)
	api.HandleFunc("/", put).Methods(http.MethodPut)
	api.HandleFunc("/", delete).Methods(http.MethodDelete)

	api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":2090", route))

}