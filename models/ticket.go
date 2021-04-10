package models

import "errors"

type Ticket struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	User   string `json:"user"`
	Status string `json:"status"`
}
type Tickets []Ticket

var tickets = make(map[int]Ticket)

func SetDefaultTicket() {
	ticket := Ticket{Id: 1, Title: "Test", User: "Mauricio", Status: "abierta"}
	tickets[ticket.Id] = ticket
}

func GetTickets() Tickets {
	list := Tickets{}
	for _, ticket := range tickets {
		list = append(list, ticket)
	}
	return list
}

func GetTicket(ticketId int) (Ticket, error) {
	if ticket, ok := tickets[ticketId]; ok {
		return ticket, nil
	}
	return Ticket{}, errors.New("El ticket no se encuentra")
}
func SaveTicket(ticket Ticket) Ticket {
	ticket.Id = len(tickets) + 1
	tickets[ticket.Id] = ticket
	return ticket
}

func UpdateTicket(ticket Ticket, title, user, status string) Ticket {
	ticket.Title = title
	ticket.User = user
	ticket.Status = status
	tickets[ticket.Id] = ticket
	return ticket
}
