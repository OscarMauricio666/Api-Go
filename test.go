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

	ticket := models.CreateTicket("Prueba666", "óscar Mauricio Rondón fajardo", "Activo")
	fmt.Println(ticket)
	// ticket.Title = "change the title"
	// ticket.Save()
	// ticket.Delete()
	models.CloseConection()
}
