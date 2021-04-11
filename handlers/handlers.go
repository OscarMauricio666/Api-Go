package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
)

func GetTickets(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetTickets())
}

func GetTicket(w http.ResponseWriter, r *http.Request) {

	if ticket, err := getTicketByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		models.SendData(w, ticket)
	}
}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	ticket := models.Ticket{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&ticket); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		models.SendData(w, models.SaveTicket(ticket))
	}
}
func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	ticket, err := getTicketByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}
	ticketResponse := models.Ticket{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&ticketResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

	ticket = models.UpdateTicket(ticket, ticketResponse.Title, ticketResponse.User, ticketResponse.Status)
	models.SendData(w, ticket)
}
func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	if ticket, err := getTicketByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		models.DeleteTicket(ticket.Id)
		models.SendNoContent(w)
	}
}

func getTicketByRequest(r *http.Request) (models.Ticket, error) {
	vars := mux.Vars(r)
	ticketId, _ := strconv.Atoi(vars["id"])

	if ticket, err := models.GetTicket(ticketId); err != nil {
		return ticket, err
	} else {
		return ticket, nil
	}
}
