package View

/*el req es un map , no es solo un array*/
type DistanceMessageRequest struct {
	Message  []string `json:"message,omitempty"`
	Distance float32  `json:"distance,omitempty"`
}

/*hacerle gets y sets*/

func (dm *DistanceMessageRequest) SetDistanceMessageRequest(message []string, distance float32) {
	dm.Message = message
	dm.Distance = distance
}

func (dm DistanceMessageRequest) GetDistanceMessageRequest() ([]string, float32) {
	return dm.Message, dm.Distance
}
