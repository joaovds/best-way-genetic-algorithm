package core

import "github.com/stretchr/testify/mock"

type mockDistanceCalculator struct {
	mock.Mock
}

func NewMockDistanceCalculator() *mockDistanceCalculator {
	return &mockDistanceCalculator{}
}

func (m *mockDistanceCalculator) CalculateDistance(from, to *Gene) float64 {
	args := m.Called(from, to)
	return args.Get(0).(float64)
}

var mockGetCacheInstanceFn = func() *Cache {
	return new(Cache)
}
