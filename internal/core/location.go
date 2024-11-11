package core

type Location struct {
	Address    string
	ID         int
	isStarting bool
}

func NewLocation(id int, address string, isStarting bool) *Location {
	return &Location{address, id, isStarting}
}
