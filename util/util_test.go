package util

import (
	"ejercicio-triangulos/models"
	"testing"
)

func TestReadMakePoint(t *testing.T) {
	fileName := "datostest1.txt"
	expected := []models.Point{
		{
			X: -2,
			Y: 3,
		},
		{
			X: 1,
			Y: 3,
		},
		{
			X: 1,
			Y: 1,
		},
		{
			X: -2,
			Y: 1,
		},
	}

	puntos, _ := ReadMakePoinst(fileName)

	if len(*puntos) != len(expected) {
		t.Errorf("el tamaño del slice obtenido no coicide con el esperado, obtuve: %d, experaba: %d.", len(*puntos), len(expected))
	} else {
		for i := range *puntos {
			if (*puntos)[i] != expected[i] {
				t.Errorf("hay un punto que no coinciden, obtuve: %v, experaba: %v.", (*puntos)[i], expected[i])
				break
			}
		}
	}

}

func TestMakeTrinagles(t *testing.T) {
	puntos := []models.Point{
		{
			X: -2,
			Y: 3,
		},
		{
			X: 1,
			Y: 3,
		},
		{
			X: 1,
			Y: 1,
		},
		{
			X: -2,
			Y: 1,
		},
	}
	expected := []models.Triangle{
		{
			A: &puntos[0],
			B: &puntos[1],
			C: &puntos[2],
		},
		{
			A: &puntos[0],
			B: &puntos[1],
			C: &puntos[3],
		},
		{
			A: &puntos[0],
			B: &puntos[2],
			C: &puntos[3],
		},
		{
			A: &puntos[1],
			B: &puntos[2],
			C: &puntos[3],
		},
	}

	triangulos := MakeTrinagles(&puntos)

	if len(*triangulos) != len(expected) {
		t.Errorf("el tamaño del slice obtenido no coicide con el esperado, obtuve: %d, experaba: %d.", len(*triangulos), len(expected))
	} else {
		for i := range *triangulos {
			if (*triangulos)[i] != expected[i] {
				t.Error("hay un triangulo que no coincide, obtuve: ", (*triangulos)[i], ", experaba: ", expected[i], ".")
			}
		}
	}
}

func TestOrientationPoints(t *testing.T) {
	puntos := []models.Point{
		{
			X: -2,
			Y: 3,
		},
		{
			X: 1,
			Y: 3,
		},
		{
			X: 1,
			Y: 1,
		},
	}
	var expected float32 = -6

	result := OrientationPoints(&puntos[0], &puntos[1], &puntos[2])

	if result != expected {
		t.Errorf("la orientacion que obtube no coicide con el esperado, obtuve: %0.2f, esperaba: %0.2f", result, expected)
	}
}

func TestTrianglePoint(t *testing.T) {
	puntosSinPuntosInterior := []models.Point{{X: -2, Y: 3}, {X: 1, Y: 3}, {X: 1, Y: 1}, {X: -2, Y: 1}}
	puntosConPuntosInterior := []models.Point{{X: -1.78, Y: 1.61}, {X: 1.62, Y: -0.69}, {X: 0.56, Y: 3.09}, {X: -0.48, Y: 1.73}}
	testCases := []struct {
		Name                 string
		Poinst               *[]models.Point
		Triangles            *[]models.Triangle
		ExpectedWithPoint    *[]models.Triangle
		ExpectedWithOutPoint *[]models.Triangle
	}{
		{
			Name:   "Triangulos sin ningun punto interior",
			Poinst: &puntosSinPuntosInterior,
			Triangles: &[]models.Triangle{
				{
					A: &puntosSinPuntosInterior[0],
					B: &puntosSinPuntosInterior[1],
					C: &puntosSinPuntosInterior[2],
				},
				{
					A: &puntosSinPuntosInterior[0],
					B: &puntosSinPuntosInterior[1],
					C: &puntosSinPuntosInterior[3],
				},
				{
					A: &puntosSinPuntosInterior[0],
					B: &puntosSinPuntosInterior[2],
					C: &puntosSinPuntosInterior[3],
				},
				{
					A: &puntosSinPuntosInterior[1],
					B: &puntosSinPuntosInterior[2],
					C: &puntosSinPuntosInterior[3],
				},
			},
			ExpectedWithPoint: new([]models.Triangle),
			ExpectedWithOutPoint: &[]models.Triangle{
				{
					A: &puntosSinPuntosInterior[0],
					B: &puntosSinPuntosInterior[1],
					C: &puntosSinPuntosInterior[2],
				},
				{
					A: &puntosSinPuntosInterior[0],
					B: &puntosSinPuntosInterior[1],
					C: &puntosSinPuntosInterior[3],
				},
				{
					A: &puntosSinPuntosInterior[0],
					B: &puntosSinPuntosInterior[2],
					C: &puntosSinPuntosInterior[3],
				},
				{
					A: &puntosSinPuntosInterior[1],
					B: &puntosSinPuntosInterior[2],
					C: &puntosSinPuntosInterior[3],
				},
			},
		},
		{
			Name:   "Triangulos con uno solo con un punto interior",
			Poinst: &puntosConPuntosInterior,
			Triangles: &[]models.Triangle{
				{
					A: &puntosConPuntosInterior[0],
					B: &puntosConPuntosInterior[1],
					C: &puntosConPuntosInterior[2],
				},
				{
					A: &puntosConPuntosInterior[0],
					B: &puntosConPuntosInterior[1],
					C: &puntosConPuntosInterior[3],
				},
				{
					A: &puntosConPuntosInterior[0],
					B: &puntosConPuntosInterior[2],
					C: &puntosConPuntosInterior[3],
				},
				{
					A: &puntosConPuntosInterior[1],
					B: &puntosConPuntosInterior[2],
					C: &puntosConPuntosInterior[3],
				},
			},
			ExpectedWithPoint: &[]models.Triangle{
				{
					A: &puntosConPuntosInterior[0],
					B: &puntosConPuntosInterior[1],
					C: &puntosConPuntosInterior[2],
				},
			},
			ExpectedWithOutPoint: &[]models.Triangle{
				{
					A: &puntosConPuntosInterior[0],
					B: &puntosConPuntosInterior[1],
					C: &puntosConPuntosInterior[3],
				},
				{
					A: &puntosConPuntosInterior[0],
					B: &puntosConPuntosInterior[2],
					C: &puntosConPuntosInterior[3],
				},
				{
					A: &puntosConPuntosInterior[1],
					B: &puntosConPuntosInterior[2],
					C: &puntosConPuntosInterior[3],
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			resultTriConPunto, resultTriSinPunto := TrianglePoint(tc.Triangles, tc.Poinst)
			for i := range *resultTriConPunto {
				if (*resultTriConPunto)[i] != (*tc.ExpectedWithPoint)[i] {
					t.Errorf("los triangulos con punto interior no coincide con el esperado, obtenido: %v, esperaba: %v.", (*resultTriConPunto)[i], (*tc.ExpectedWithPoint)[i])
				}
			}
			for i := range *resultTriSinPunto {
				if (*resultTriSinPunto)[i] != (*tc.ExpectedWithOutPoint)[i] {
					t.Errorf("los triangulos sin punto interior no coincide con el esperado, obtenido: %v, esperaba: %v.", (*resultTriSinPunto)[i], (*tc.ExpectedWithOutPoint)[i])
				}
			}
		})
	}
}

func TestSuperficieMayor(t *testing.T) {
	puntos := []models.Point{{X: -1.78, Y: 1.61}, {X: 1.62, Y: -0.69}, {X: 0.56, Y: 3.09}, {X: -0.48, Y: 1.73}}
	triangulos := []models.Triangle{
		{
			A: &puntos[0],
			B: &puntos[1],
			C: &puntos[2],
		},
		{
			A: &puntos[0],
			B: &puntos[1],
			C: &puntos[3],
		},
		{
			A: &puntos[0],
			B: &puntos[2],
			C: &puntos[3],
		},
		{
			A: &puntos[1],
			B: &puntos[2],
			C: &puntos[3],
		},
	}
	expected := triangulos[0]

	result := SuperficieMayor(&triangulos)

	if *result != expected {
		t.Errorf("el triangulo de mayor superficie coincide con el esperado, obtenido: %v, esperaba: %v.", *result, expected)
	}
}
