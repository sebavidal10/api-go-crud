package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sebavidal10/api-go-crud/pkg/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/events", handlers.GettAllEvents).Methods(http.MethodGet)
	router.HandleFunc("/api/events", handlers.AddEvent).Methods(http.MethodPost)
	router.HandleFunc("/api/events/{id}", handlers.GetEvent).Methods(http.MethodGet)
	router.HandleFunc("/api/events/{id}", handlers.UpdateEvent).Methods(http.MethodPut)
	router.HandleFunc("/api/events/{id}", handlers.DeleteEvent).Methods(http.MethodDelete)

	log.Println("Listening on port 4000")
	http.ListenAndServe(":4000", router)
}
