package test

import (
	"testing"

	"../models"
)

func TestConnection(t *testing.T) {
	connection := models.GetConnection()
	if connection == nil {
		t.Error("No es posible conectarse")
	} else {
		println("Pasa la prueba de conección a la base de datos")
	}
}
