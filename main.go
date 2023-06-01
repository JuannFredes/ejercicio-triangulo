package main

import (
	"fmt"
	"os"

	"github.com/juanjgfredes/ejercicio-triangulo.git/util"
)

func main() {
	//Leo los datos.txt y si hay un error lo miestro y termino el programa
	nameFile := "datos.txt"
	points, err := util.ReadMakePoinst(nameFile)
	if err != nil {
		fmt.Println("ocurrio un error: " + err.Error())
		os.Exit(1)
	}
	//Formo los triangulos con los puntos leidos
	triangles := util.MakeTrinagles(points)
	//Obtengo los triangulos con y sin punto interior
	trianglesPointInside, trianglesPointOutside := util.TrianglePoint(triangles, points)
	//Busco el de mayor superficie en ambos caso
	trianguloSuperficieMayorPuntoInterior := util.SuperficieMayor(trianglesPointInside)
	trianguloSuperficieMayorSinPuntoInterior := util.SuperficieMayor(trianglesPointOutside)
	//Imprimo el resultado
	fmt.Printf("El triangulo, con almenos un punto interior, de mayor superficie es: %v \n"+
		"El triangulo, sin ningun punto interior, de mayor superficie es: %v \n", *trianguloSuperficieMayorPuntoInterior, *trianguloSuperficieMayorSinPuntoInterior)
}
