package View

import (
	"quasar/Model"
)

/*el req es un map , no es solo un array*/
type SatellitesRequest struct {
	Satellites map[string][]Model.Satellite `json:"satellites,omitempty"`
}

/*hacerle gets y sets*/

func (s *SatellitesRequest) SetSatellites(satellites map[string][]Model.Satellite) {
	s.Satellites = satellites
}

func (s SatellitesRequest) GetSatellites() map[string][]Model.Satellite {
	return s.Satellites
}
