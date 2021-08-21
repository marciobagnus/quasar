package Model

type Satellite struct {
	Name     string   `json:"Name,omitempty"`
	Distance float64  `json:"Distance,omitempty"`
	Message  []string `json:"Message,omitempty"`
}
