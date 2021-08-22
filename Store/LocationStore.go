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

//newLocationStore crea un almacenamiento de locations
func NewLocationsStore() *LocationsStore {
	return &LocationsStore{
		/*creo el store con los datos ya. */
		store: locations,
		//store: map[string]Model.Location{},
	}

}

//SaveLocations guarda locations
func (l *LocationsStore) SaveLocations(locations map[string]Model.Location) {

	l.Lock()
	l.store = locations
	defer l.Unlock()
}

func (l *LocationsStore) FetchLocationByName(name string) (location Model.Location, err error) {

	l.Lock()

	//Forma de buscar sacada de stackoverflow
	if val, ok := l.store[strings.ToUpper(name)]; ok {
		l.Unlock()
		return val, nil
	}
	//otra forma de buscar
	/*i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}*/
	return location, errors.New("The Satelite: " + name + " doesn't exist")
}
