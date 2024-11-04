package dto

type Location struct {
	ID int
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
}

func NewLocation(x, y float64) *Location {
	return &Location{X: x, Y: y}
}
