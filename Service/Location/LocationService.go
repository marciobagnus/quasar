package Location

import (
	"errors"
	"math"
	"quasar/Model"
	"quasar/Store"
)

/*
Obtiene la localizacion del emisor por trilateracion
Explicacion de trilateracion
https://math.stackexchange.com/questions/884807/find-x-location-using-3-known-x-y-location-using-trilateration

Explicacion y algoritmo en python
https://www.101computing.net/cell-phone-trilateration-algorithm/

Se podrian tratar el tema de los muy grandes libreria: https://github.com/shopspring/decimal
*/
func GetLocation(mapSatellites map[string][]Model.Satellite) (location Model.Location, err error) {

	var satellites []Model.Satellite

	/*Recorre el map que tiene una sola key y adentro de esa key estan los satelites*/
	for _, v := range mapSatellites {
		satellites = v
	}

	/*Crea el store de location*/
	locationStore := Store.NewLocationsStore()

	if len(satellites) < 3 {
		return location, errors.New("incorrect number of satellites.")
	}

	//Voy desarmando los datos de cada satelite
	s1 := satellites[0]
	r1 := s1.Distance
	//Lo declaro asi porque del otro lado tambien me devuelve error
	location1, err := locationStore.FetchLocationBySatelliteName(s1.Name)
	x1 := location1.X
	y1 := location1.Y

	s2 := satellites[1]
	r2 := s2.Distance
	location2, err := locationStore.FetchLocationBySatelliteName(s2.Name)
	x2 := location2.X
	y2 := location2.Y

	s3 := satellites[2]
	r3 := s3.Distance
	location3, err := locationStore.FetchLocationBySatelliteName(s3.Name)
	x3 := location3.X
	y3 := location3.Y

	A := calculateDifference(x2, x1)
	B := calculateDifference(y2, y1)
	C := calculateMatrix(r1, r2, x1, x2, y1, y2)
	//C := r1*r1 - r2*r2 - x1*x1 + x2*x2 - y1*y1 + y2*y2
	D := calculateDifference(x3, x2)
	E := calculateDifference(y3, y2)
	F := calculateMatrix(r2, r3, x2, x3, y2, y3)
	//F := r2*r2 - r3*r3 - x2*x2 + x3*x3 - y2*y2 + y3*y3
	x := calculateAxisPoint(C, E, F, B, E, A, B, D)
	//x := (C*E - F*B) / (E*A - B*D)
	y := calculateAxisPoint(C, D, A, F, B, D, A, E)
	//y := (C*D - A*F) / (B*D - A*E)

	location.X = x
	location.Y = y

	return location, nil

}

/*Calcula la diferencia entre el doble de cada variable*/
func calculateDifference(a float32, b float32) float32 {
	return 2*a - 2*b
}

/*Hace el calculo con la diferencia de los cuadrados expandidos */
func calculateMatrix(a float32, b float32, c float32, d float32, e float32, f float32) float32 {
	return a*a - b*b - c*c + d*d - e*e + f*f
}

/*Calcula el punto x y en el plano*/
func calculateAxisPoint(v1 float32, v2 float32, v3 float32, v4 float32, v5 float32, v6 float32, v7 float32, v8 float32) float32 {
	value := (v1*v2 - v3*v4) / (v5*v6 - v7*v8)
	valueCast := float64(value)
	valueRound := math.Round(valueCast*100) / 100
	return float32(valueRound)
}
