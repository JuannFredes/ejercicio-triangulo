package models

import (
	"fmt"
	"math"
)

/*
Modelo de los triagulos que se formaran con los puntos
*/

type Triangle struct {
	A *Point
	B *Point
	C *Point
}

func (t Triangle) String() string {
	if t.A == nil && t.B == nil && t.C == nil {
		return "Triangulo nulo"
	}
	return fmt.Sprintf("\n{ \n"+
		"   A: %v,\n"+
		"   B: %v,\n"+
		"   C: %v\n"+
		"}", t.A, t.B, t.C)
}

/*
Este metodo calcula y retorna la superficie del triangulo
*/

func (t *Triangle) Superficie() float64 {
	ladoA := math.Sqrt(math.Pow(float64(t.B.X)-float64(t.A.X), 2) + math.Pow(float64(t.B.Y)-float64(t.A.Y), 2))
	ladoB := math.Sqrt(math.Pow(float64(t.C.X)-float64(t.B.X), 2) + math.Pow(float64(t.C.Y)-float64(t.B.Y), 2))
	ladoC := math.Sqrt(math.Pow(float64(t.C.X)-float64(t.A.X), 2) + math.Pow(float64(t.C.Y)-float64(t.A.Y), 2))

	return ladoA + ladoB + ladoC
}
