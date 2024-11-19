package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCoreLocation(t *testing.T) {
	t.Run("Should convert LocationRequest to core Locations with a single starting point", func(t *testing.T) {
		locationRequest := LocationRequest{
			Locations: []Location{
				{Address: "Start", IsStarting: true},
				{Address: "Loc1", IsStarting: false},
				{Address: "Loc2", IsStarting: false},
			},
		}

		startingPoint, locations, err := locationRequest.ToCoreLocation()

		assert.NoError(t, err)
		assert.NotNil(t, startingPoint)
		assert.Equal(t, "Start", startingPoint.Address)
		assert.Len(t, locations, 2)
		assert.Equal(t, "Loc1", locations[0].Address)
		assert.Equal(t, "Loc2", locations[1].Address)
	})

	t.Run("Should use the first location as starting point if no starting point is specified", func(t *testing.T) {
		locationRequest := LocationRequest{
			Locations: []Location{
				{Address: "Loc1", IsStarting: false},
				{Address: "Loc2", IsStarting: false},
				{Address: "Loc3", IsStarting: false},
			},
		}

		startingPoint, locations, err := locationRequest.ToCoreLocation()

		assert.NoError(t, err)
		assert.NotNil(t, startingPoint)
		assert.Equal(t, "Loc1", startingPoint.Address)
		assert.Len(t, locations, 2)
		assert.Equal(t, "Loc2", locations[0].Address)
		assert.Equal(t, "Loc3", locations[1].Address)
	})

	t.Run("Should ignore additional starting points and only use the first one", func(t *testing.T) {
		locationRequest := LocationRequest{
			Locations: []Location{
				{Address: "Start", IsStarting: true},
				{Address: "Loc1", IsStarting: false},
				{Address: "Another Start", IsStarting: true},
				{Address: "Loc2", IsStarting: false},
			},
		}

		startingPoint, locations, err := locationRequest.ToCoreLocation()

		assert.NoError(t, err)
		assert.NotNil(t, startingPoint)
		assert.Equal(t, "Start", startingPoint.Address)
		assert.Len(t, locations, 3)
		assert.Equal(t, "Loc1", locations[0].Address)
		assert.Equal(t, "Another Start", locations[1].Address)
		assert.Equal(t, "Loc2", locations[2].Address)
	})

	t.Run("Should return error for empty LocationRequest", func(t *testing.T) {
		locationRequest := LocationRequest{}

		startingPoint, locations, err := locationRequest.ToCoreLocation()

		assert.Error(t, err)
		assert.Nil(t, startingPoint)
		assert.Nil(t, locations)
	})

	t.Run("Should return error for LocationRequest with only one location", func(t *testing.T) {
		locationRequest := LocationRequest{
			Locations: []Location{
				{Address: "Start", IsStarting: true},
			},
		}

		startingPoint, locations, err := locationRequest.ToCoreLocation()

		assert.Error(t, err)
		assert.Nil(t, startingPoint)
		assert.Nil(t, locations)
	})

	t.Run("Should return error for LocationRequest with only two locations", func(t *testing.T) {
		locationRequest := LocationRequest{
			Locations: []Location{
				{Address: "Start", IsStarting: true},
				{Address: "Loc1", IsStarting: false},
			},
		}

		startingPoint, locations, err := locationRequest.ToCoreLocation()

		assert.Error(t, err)
		assert.Nil(t, startingPoint)
		assert.Nil(t, locations)
	})
}
