package models

import (
	"math"
	"testing"
)

func TestSuperficie(t *testing.T) {
	triangulo := Triangle{
		A: &Point{0, 4},
		B: &Point{4, 0},
		C: &Point{-4, 0},
	}
	var expected float64 = 19.313708
	result := math.Round(triangulo.Superficie()*1e6) / 1e6

	if result != expected {
		t.Errorf("la superficie no coincide con la esperaba, obtenida: %v, esperaba: %v.", result, expected)
	}
}
