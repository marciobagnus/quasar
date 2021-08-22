package Store

import (
	"quasar/Model"
	"sync"
)

//satellitesStore es el almacenamiento en si
type SatellitesStore struct {
	sync.Mutex
	store map[string][]Model.Satellite
}

//newSatelliteStore crea un almacenamiento de satelites
func NewSatelliteStore() *SatellitesStore {
	return &SatellitesStore{
		store: map[string][]Model.Satellite{},
	}

}

//SaveSatellites guarda los satelites
func (s *SatellitesStore) SaveSatellites(satellites map[string][]Model.Satellite) {

	s.Lock()
	s.store = satellites
	defer s.Unlock()
}
