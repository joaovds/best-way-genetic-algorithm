package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGene(t *testing.T) {
	gene := NewGene(1, "any_address")
	assert.Equal(t, 1, gene.GetID())
	assert.Equal(t, "any_address", gene.Address)
	assert.Equal(t, 0.0, gene.Distance)
}

func TestCalculateDistanceToDestination(t *testing.T) {
	t.Run("should return 0 if destinationID is equal to fromID", func(t *testing.T) {
		for i := range 10 {
			from := NewGene(i, "any_address")
			destination := from
			calculatorMock := NewMockDistanceCalculator()
			mockCache := mockGetCacheInstanceFn()
			result := from.CalculateDistanceToDestination(destination, calculatorMock, mockCache)
			assert.Empty(t, result)
			calculatorMock.AssertNotCalled(t, "CalculateDistance")
		}
	})

	t.Run("should calculate distance if destinationID is different from fromID", func(t *testing.T) {
		from := NewGene(1, "any_address")
		destination := NewGene(2, "another_address")
		calculatorMock := NewMockDistanceCalculator()
		mockCache := mockGetCacheInstanceFn()
		calculatorMock.On("CalculateDistance", from, destination).Return(17.77)

		result := from.CalculateDistanceToDestination(destination, calculatorMock, mockCache)
		assert.Equal(t, 17.77, result)
		calculatorMock.AssertCalled(t, "CalculateDistance", from, destination)
	})

	t.Run("should use cache if distance is already calculated", func(t *testing.T) {
		mockCache := mockGetCacheInstanceFn()
		for i := range 10 {
			from := NewGene(i, "any_address")
			destination := NewGene(i+1, "another_address")
			calculatorMock := &mockDistanceCalculator{}
			calculatorMock.On("CalculateDistance", from, destination).Return(22.876)

			result := from.CalculateDistanceToDestination(destination, calculatorMock, mockCache)
			assert.Equal(t, 22.876, result)
			calculatorMock.AssertCalled(t, "CalculateDistance", from, destination)

			result = from.CalculateDistanceToDestination(destination, calculatorMock, mockCache)
			calculatorMock.AssertNotCalled(t, "CalculateDistance")
			calculatorMock.AssertNumberOfCalls(t, "CalculateDistance", 1)
		}
	})
}
