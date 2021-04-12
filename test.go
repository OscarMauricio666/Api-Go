package main

import (
	"fmt"

	"./models"
)

func main() {
	models.CreateConection()
	// models.Ping()
	models.CreateTables()
	models.CreateTicket("Prueba666", "óscar Mauricio Rondón fajardo", "Activo")
	models.CreateTicket("333", "óscar Mauricio Rondón fajardo", "Activo")
	models.CreateTicket("999", "óscar Mauricio Rondón fajardo", "Inactivo")

	ticket666 := models.GetTicket(1)
	fmt.Println(ticket666)

	ticket999 := models.GetTickets()
	fmt.Println(ticket999)

	ticket := models.CreateTicket("Prueba666", "óscar Mauricio Rondón fajardo", "Activo")
	fmt.Println(ticket)
	// ticket.Title = "change the title"
	// ticket.Save()
	// ticket.Delete()
	models.CloseConection()
}
