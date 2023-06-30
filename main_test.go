package main

import "testing"

func TestValidar(t *testing.T) {
	array1 := [...]int {10, 33, 22, 654, 34, 988, 872, 104, 358, 3}

	result := buscarMinimo(array1)
	if result != 3 {
		t.Errorf("Fallo test, se obtuvo: %d, se esperaba: %d.", result, 3)
	}

	result = buscarMaximo(array1)
	if result != 988 {
		t.Errorf("Fallo test, se obtuvo: %d, se esperaba: %d.", result, 988)
	}

	resultFloat := calcularPromedio(array1)
	if resultFloat != 306.8 {
		t.Errorf("Fallo test, se obtuvo: %f, se esperaba: %f.", resultFloat, 306.8)
	}

	resultFloat = calcularSumatoria(array1)
	if resultFloat != 3068.0 {
		t.Errorf("Fallo test, se obtuvo: %f, se esperaba: %f.", resultFloat, 3068.0)
	}
}
