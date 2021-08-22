package Service

import (
	"errors"
	"quasar/Model"
	"quasar/Store"
)

func GetLocation(mapSatellites map[string][]Model.Satellite) (location Model.Location, err error) {

	var satellites []Model.Satellite

	/*Creo el store de location desde aca porue locoy a necesitar*/

	locationStore := Store.NewLocationsStore()

	/*recoro el map que tiene una sola key y adentro de esa key estan los satelites*/

	for _, v := range mapSatellites {
		satellites = v
	}

	if len(satellites) < 3 {
		return location, errors.New("Incorrect number of satellites.")
	}

	//Voy desarmando los datos de cada satelite

	s1 := satellites[0]
	r1 := s1.Distance
	//Lo declaro asi porque del otro lado tambien me devuelve error
	location1, err := locationStore.FetchLocationByName(s1.Name)
	x1 := location1.X
	y1 := location1.Y

	s2 := satellites[1]
	r2 := s2.Distance
	location2, err := locationStore.FetchLocationByName(s2.Name)
	x2 := location2.X
	y2 := location2.Y

	s3 := satellites[2]
	r3 := s3.Distance
	location3, err := locationStore.FetchLocationByName(s3.Name)
	x3 := location3.X
	y3 := location3.Y

	A := 2*x2 - 2*x1
	B := 2*y2 - 2*y1
	C := r1*r1 - r2*r2 - x1*x1 + x2*x2 - y1*y1 + y2*y2
	D := 2*x3 - 2*x2
	E := 2*y3 - 2*y2
	F := r2*r2 - r3*r3 - x2*x2 + x3*x3 - y2*y2 + y3*y3

	x := (C*E - F*B) / (E*A - B*D)
	y := (C*D - A*F) / (B*D - A*E)

	location.X = x
	location.Y = y

	return location, nil

}
