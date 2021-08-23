package Store

import (
	"quasar/Model"
	"strings"
	"sync"
)

//satellitesStore es el almacenamiento en si
type SatellitesStore struct {
	sync.Mutex
	store map[string]Model.Satellite
}

/*inicializoel mapa para restringir los nombres */
var satellites = map[string]Model.Satellite{
	"KENOBI":    {},
	"SKYWALKER": {},
	"SATO":      {},
}

//newSatelliteStore crea un almacenamiento de satelites
func NewSatellitesStore() *SatellitesStore {
	return &SatellitesStore{
		store: satellites,
	}

}

/*agregar satelites de a uno*/

func (s *SatellitesStore) SaveSatellite(satellite Model.Satellite, name string) (status string) {

	/*primero obtengo el satelite por el nombre*/
	//var sat Model.Satellite
	s.Lock()
	//Forma de buscar sacada de stackoverflow
	/*si ya existe eliminarlo*/
	if val, ok := s.store[strings.ToUpper(name)]; ok {
		s.Unlock()
		//sat = val
		delete(s.store, strings.ToUpper(val.Name))
		s.store[strings.ToUpper(name)] = satellite
		return "Satellite saved"
	}
	return "The satellite name is incorrect"

}

/*
func (s *SatellitesStore) GetSatellites() (satellites map[string]Model.Satellite) {

	s.Lock()
	if len(s.store) > 0 {
		s.Unlock()
		return s.store
	}
	s.Unlock()

	return satellites

}*/

func (s *SatellitesStore) GetSatellites() (satellites map[string][]Model.Satellite) {

	//Tengo que transformar el mapa de un map[string]Model.Satellite a map[string][]Model.Satellites */
	//Primero obtengo el array de satelites*/
	var satellitesArray []Model.Satellite

	s.Lock()
	for _, sat := range s.store {
		satellitesArray = append(satellitesArray, sat)
	}
	s.Unlock()

	/*Luego creo un mapa de arrays*/
	//var satellites map[string][]Model.Satellite
	//satellites = make(map[string][]Model.Satellite)
	//Aca recorro el el array anterior para appendearle cada satelite a la unica key "satellites"

	satNew := make(map[string][]Model.Satellite)
	satNew["satellites"] = append(satellites["satellites"], satellitesArray...)

	satellites = satNew

	return satellites

}
