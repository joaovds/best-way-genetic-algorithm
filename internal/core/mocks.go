package core

import "github.com/stretchr/testify/mock"

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

type MockCrossover struct {
	mock.Mock
}

func NewMockCrossover() *MockCrossover {
	return &MockCrossover{}
}

func (m *MockCrossover) Run() [2]*Chromosome {
	args := m.Called()
	return args.Get(0).([2]*Chromosome)
}

// ----- ... -----
