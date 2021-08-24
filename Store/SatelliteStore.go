package Store

import (
	"quasar/Model"
	"strings"
	"sync"
)

//Representa el almacenamiento en si
type SatellitesStore struct {
	sync.Mutex
	store map[string]Model.Satellite
}

/*Inicializo el mapa para restringir los nombres */
var satellites = map[string]Model.Satellite{
	"KENOBI":    {},
	"SKYWALKER": {},
	"SATO":      {},
}

//Crea un almacenamiento de satelites
func NewSatellitesStore() *SatellitesStore {
	return &SatellitesStore{
		store: satellites,
	}
}

//Guarda un satelites*/
func (s *SatellitesStore) SaveSatellite(satellite Model.Satellite, name string) (status string) {

	s.Lock()
	//Forma de buscar en un map sacada de stackoverflow. /*Primero obtengo el satelite por el nombre*/
	/*si ya existe eliminarlo y volverlo a agregar*/
	if val, ok := s.store[strings.ToUpper(name)]; ok {
		s.Unlock()
		delete(s.store, strings.ToUpper(val.Name))
		s.store[strings.ToUpper(name)] = satellite
		return "Satellite saved"
	}
	return "The satellite name is incorrect"

}

func (s *SatellitesStore) GetSatellites() (satellites map[string][]Model.Satellite) {

	/*Tengo que transformar el mapa de un map[string]Model.Satellite a map[string][]Model.Satellites
	Primero obtengo el array de satelites*/
	var satellitesArray []Model.Satellite

	s.Lock()
	for _, sat := range s.store {
		satellitesArray = append(satellitesArray, sat)
	}
	s.Unlock()

	/*Creo un mapa de arrays
	Recorro el el array anterior para appendearle cada satelite a la unica key "satellites"*/
	satNew := make(map[string][]Model.Satellite)
	satNew["satellites"] = append(satellites["satellites"], satellitesArray...)
	satellites = satNew
	return satellites

}
