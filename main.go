package main

import (
	"log"
	"net/http"

	"./handlers"
	"./models"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	models.SetDefaultTicket()
	mux.HandleFunc("/api/tickets/", handlers.GetTickets).Methods("GET")
	mux.HandleFunc("/api/tickets/{id:[0-9]+}", handlers.GetTicket).Methods("GET")
	mux.HandleFunc("/api/tickets/", handlers.CreateTicket).Methods("POST")
	mux.HandleFunc("/api/tickets/{id:[0-9]+}", handlers.UpdateTicket).Methods("PUT")
	mux.HandleFunc("/api/tickets/{id:[0-9]+}", handlers.DeleteTicket).Methods("DELETE")
	log.Println("El servidor esta Arriba puerto 8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
