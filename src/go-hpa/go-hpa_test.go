package main

import "testing"

func TestDosqrt(t *testing.T) {
	total := dosqrt(9);
	if total != 3.0 {
		t.Errorf("Saída incorreta, recebi: %.1f, esperado: %s.", total, "3.0")
	}
}
