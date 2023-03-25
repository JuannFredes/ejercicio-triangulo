package models

import "fmt"

/*
Hago el modelo de los puntos (x,y)
*/

type Point struct {
	X float32
	Y float32
}

func (p *Point) String() string {
	return fmt.Sprintf("{X: %f, Y: %f}", p.X, p.Y)
}
