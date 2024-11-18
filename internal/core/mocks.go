package core

import "github.com/stretchr/testify/mock"

type MockDistanceCalculator struct {
	mock.Mock
}

func NewMockDistanceCalculator() *MockDistanceCalculator {
	return &MockDistanceCalculator{}
}

func (m *MockDistanceCalculator) CalculateDistance(from, to *Gene) float64 {
	args := m.Called(from, to)
	return args.Get(0).(float64)
}

// ----- ... -----

var MockGetCacheInstanceFn = func() *Cache {
	return new(Cache)
}

// ----- ... -----

type MockSelection struct {
	mock.Mock
}

func NewMockSelection() *MockSelection {
	return &MockSelection{}
}

func (m *MockSelection) Select(p *Population) *Chromosome {
	args := m.Called(p)
	return args.Get(0).(*Chromosome)
}

// ----- ... -----
