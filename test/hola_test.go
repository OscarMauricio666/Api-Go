package test

import "testing"

func TestHolaMundo(t *testing.T) {
	str := "hola Mundo"
	if str != "hola Mundo" {
		t.Error("No es Posible saludar a los usuarios", nil)
	} else {
		println("Pasa prueba")
	}
}
