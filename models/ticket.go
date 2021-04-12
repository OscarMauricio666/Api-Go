package models

type Ticket struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	User   string `json:"user"`
	Status string `json:"status"`
}

const ticketsSchema string = `CREATE TABLE tickets(
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(30) NOT NULL,
    user VARCHAR(66) NOT NULL,
    status VARCHAR(66) NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

type Tickets []Ticket

func NewTicket(title, user, status string) *Ticket {
	ticket := &Ticket{Title: title, User: user, Status: status}
	return ticket
}
func CreateTicket(title, user, status string) *Ticket {
	ticket := NewTicket(title, user, status)
	ticket.Save()
	return ticket

}
func (this *Ticket) Save() {
	if this.Id == 0 {
		this.insert()
	} else {
		this.update()
	}
}
func (this *Ticket) insert() {
	sql := "INSERT tickets SET title=?, user=?, status=?" //? evita sql injection
	result, _ := Exec(sql, this.Title, this.User, this.Status)
	this.Id, _ = result.LastInsertId()
}
func (this *Ticket) update() {
	sql := "UPDATE tickets SET title=?, user=?, status=?" //? evita sql injection
	Exec(sql, this.Title, this.User, this.Status)
}
func (this *Ticket) Delete() {
	sql := "DELETE FROM tickets WHERE id=?"
	Exec(sql, this.Id)
}
func GetTicket(id int) *Ticket {
	ticket := NewTicket("", "", "")
	sql := "SELECT id, title, user, status FROM tickets WHERE id=?"
	rows, _ := Query(sql, id)
	for rows.Next() {
		rows.Scan(&ticket.Id, &ticket.Title, &ticket.User, &ticket.Status)
	}
	return ticket
}
func GetTickets() Tickets {
	sql := "SELECT id, title, user, status FROM tickets"
	tickets := Tickets{}
	rows, _ := Query(sql)
	for rows.Next() {
		ticket := Ticket{}
		rows.Scan(&ticket.Id, &ticket.Title, &ticket.User, &ticket.Status)
		tickets = append(tickets, ticket)
	}
	return tickets
}

// var tickets = make(map[int]Ticket)

// func SetDefaultTicket() {
// 	ticket := Ticket{Id: 1, Title: "Test", User: "Mauricio", Status: "abierta"}
// 	tickets[ticket.Id] = ticket
// }

// func GetTickets() Tickets {
// 	list := Tickets{}
// 	for _, ticket := range tickets {
// 		list = append(list, ticket)
// 	}
// 	return list
// }

// func GetTicket(ticketId int) (Ticket, error) {
// 	if ticket, ok := tickets[ticketId]; ok {
// 		return ticket, nil
// 	}
// 	return Ticket{}, errors.New("El ticket no se encuentra")
// }
// func SaveTicket(ticket Ticket) Ticket {
// 	ticket.Id = len(tickets) + 1
// 	tickets[ticket.Id] = ticket
// 	return ticket
// }

// func UpdateTicket(ticket Ticket, title, user, status string) Ticket {
// 	ticket.Title = title
// 	ticket.User = user
// 	ticket.Status = status
// 	tickets[ticket.Id] = ticket
// 	return ticket
// }
// func DeleteTicket(id int) {
// 	delete(tickets, id)

// }
