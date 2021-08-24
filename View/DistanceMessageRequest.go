package View

/*el req es un map , no es solo un array*/
type DistanceMessageRequest struct {
	Message  []string `json:"message,omitempty"`
	Distance float32  `json:"distance,omitempty"`
}
