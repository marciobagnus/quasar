package View

import (
	"quasar/Model"
)

type MessagePositionResponse struct {
	Position Model.Location `json:"position,omitempty"`
	Message  string         `json:"message,omitempty"`
}

/*hacerle gets y sets*/
