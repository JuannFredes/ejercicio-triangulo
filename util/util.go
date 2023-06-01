package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/juanjgfredes/ejercicio-triangulo.git/models"
)

/*
Esta funcion lee el archivo txt pasandole la ruta del mismo por
parametro, mientras lo lee linea por linea hice la logica necesaria
para que vaya formando los puntos con los datos de ese txt siguiendo
un determinado formato. Retorna el puntero al slice con los puntos formados
y un error en caso de suceder algun incoveniente
*/

func ReadMakePoinst(filename string) (*[]models.Point, error) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var points []models.Point
	var valueX, valueY float64

	for scanner.Scan() {
		pointRead := strings.Split(scanner.Text(), " ")

		if pointRead[0] == "0" && pointRead[1] == "0" {
			break
		}

		valueX, _ = strconv.ParseFloat(pointRead[0], 64)
		valueY, _ = strconv.ParseFloat(pointRead[1], 64)
		points = append((points), models.Point{X: float32(valueX), Y: float32(valueY)})
	}
	return &points, nil
}

/*
Esta funcion forma todos los triangulos, a partir de un puntero de
slices de puntos pasado por parametro, posibles devolviendo un puntero de
un slice con los triangulos formados con dichos puntos
*/

func MakeTrinagles(points *[]models.Point) *[]models.Triangle {
	var triangles []models.Triangle
	var triangulo models.Triangle
	a, b, c := 0, 0, 0

	for a := a; a < (len(*points) - 2); a++ {
		triangulo.A = &((*points)[a])
		b = a + 1
		for b := b; b < (len(*points) - 1); b++ {
			triangulo.B = &((*points)[b])
			c = b + 1
			for c := c; c < len(*points); c++ {
				triangulo.C = &((*points)[c])
				triangles = append(triangles, triangulo)
			}
		}
	}

	return &triangles
}

/*
Esta funcion recibe 3 puntos para sacar la orientacion del triagulo que formen dichos
puntos, y retorna el resultado del calculo
*/
func OrientationPoints(A1 *models.Point, A2 *models.Point, A3 *models.Point) float32 {
	return (A1.X-A3.X)*(A2.Y-A3.Y) - (A1.Y-A3.Y)*(A2.X-A3.X)
}

/*
Esta funcion recibe punteros que apuntan a un slices de triangulos y puntos, luego realiza
los calculos, utilizando la funcion anterior, para calcular las orientaciones del triangulo
y de los triangulos formados con el punto reemplazando algun verice del triangulo. En el
calculo de orientacion del triangulo formado con el punto que se verifica que este o no
en el interior decidi no incluir el 0 para asi que los vertices del triangulo no cuenten
como puntos interiores. La funcion retorna 2 punteros apuntando slices de triangulos, uno
con los triangulos con al menos 1 punto en su interior y otra con los triangulos sin puntos
interiores
*/

func TrianglePoint(triangles *[]models.Triangle, points *[]models.Point) (triagulosPuntoInterior, triagulosSinPuntoInterior *[]models.Triangle) {
	triagulosPuntoInterior = new([]models.Triangle)
	triagulosSinPuntoInterior = new([]models.Triangle)
	var puntoInterior bool

	for _, t := range *triangles {
		puntoInterior = false
		if OrientationPoints(t.A, t.B, t.C) >= 0 {
			for _, p := range *points {
				if OrientationPoints(t.A, t.B, &p) > 0 &&
					OrientationPoints(t.A, &p, t.C) > 0 &&
					OrientationPoints(&p, t.B, t.C) > 0 {
					*triagulosPuntoInterior = append(*triagulosPuntoInterior, t)
					puntoInterior = true
					break
				}
			}
		} else {
			for _, p := range *points {
				if OrientationPoints(t.A, t.B, &p) < 0 &&
					OrientationPoints(t.A, &p, t.C) < 0 &&
					OrientationPoints(&p, t.B, t.C) < 0 {
					*triagulosPuntoInterior = append(*triagulosPuntoInterior, t)
					puntoInterior = true
					break
				}
			}
		}
		if !puntoInterior {
			*triagulosSinPuntoInterior = append(*triagulosSinPuntoInterior, t)
		}
	}
	return
}

func SuperficieMayor(triangles *[]models.Triangle) *models.Triangle {
	var trangleSup models.Triangle
	supMax := ((*triangles)[0]).Superficie()

	for _, t := range *triangles {
		if t.Superficie() >= supMax {
			supMax = t.Superficie()
			trangleSup = t
		}
	}

	return &trangleSup
}
