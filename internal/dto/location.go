package dto

type Location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func NewLocation(x, y float64) *Location {
	return &Location{x, y}
}
