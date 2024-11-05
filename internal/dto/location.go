package dto

type Location struct {
	ID            int
	X             float64 `json:"x"`
	Y             float64 `json:"y"`
	StartingPoint bool    `json:"starting_point"`
}

func NewLocation(x, y float64, startingPoint bool) *Location {
	return &Location{X: x, Y: y, StartingPoint: startingPoint}
}
