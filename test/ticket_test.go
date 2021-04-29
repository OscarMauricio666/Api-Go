package test

import (
	"testing"

	"../models"
)

func TestNewTicket(t *testing.T) {
	ticket := models.NewTicket("title", "user", "status")
	if ticket.Title != "title" || ticket.User != "user" || ticket.Status != "status" {
		t.Error("No es posible crear el objeto ticket")
	} else {
		println("Pasa la prueba objeto ticket creado")
	}
}
func TestSave(t *testing.T) {
	ticket := models.NewTicket("title", "user", "status")
	if err := ticket.Save(); err != nil {
		t.Error("No es posible crear el usuario", err)
	} else {
		println("Pasa la prueba de Guardar el ticket")
	}
}
func TestCreateTicket(t *testing.T) {
	_, err := models.CreateTicket("title", "user", "status")
	if err != nil {
		t.Error("No es posible insertar el objeto")
	} else {
		println("Creado ticket y insertado")
	}

}
