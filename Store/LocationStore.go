package Store

import (
	"errors"
	"quasar/Model"
	"strings"
	"sync"
)

type LocationsStore struct {
	sync.Mutex
	store map[string]Model.Location
}

var locations = map[string]Model.Location{
	"KENOBI": {
		X: -500.00,
		Y: -200.00,
	},
	"SKYWALKER": {
		X: 100.00,
		Y: -100.00,
	},
	"SATO": {
		X: 500.00,
		Y: 100.00,
	},
}

//Crea un almacenamiento de locations
func NewLocationsStore() *LocationsStore {
	return &LocationsStore{
		/*Creo el store con los datos. */
		store: locations,
	}

}

func (l *LocationsStore) FetchLocationBySatelliteName(name string) (location Model.Location, err error) {

	l.Lock()
	//Forma de buscar extraida de stackoverflow
	if val, ok := l.store[strings.ToUpper(name)]; ok {
		l.Unlock()
		return val, nil
	}
	return location, errors.New("The location of Satelite: " + name + " doesn't exist")
}
