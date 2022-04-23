package main

import "testing"

func testeSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado inválido: Resultado %d. Esperado: %d", total, 30)
	}
}
