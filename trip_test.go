package gosrm

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrip(t *testing.T) {
	osrm, err := New(getOSRMAddress())
	assert.NoError(t, err)

	res, err := Trip[string](context.Background(), osrm, Request{
		Coordinates: []Coordinate{{51.3796, 35.6849}, {51.3923, 35.6882}, {51.3840, 35.6919}},
		Profile:     ProfileCar,
	})

	assert.NoError(t, err)
	assert.Equal(t, CodeOK, res.Code)
	assert.Len(t, res.Waypoints, 3)

	osrm.baseURL.Host = "invalid"
	res, err = Trip[string](context.Background(), osrm, Request{
		Coordinates: []Coordinate{{51.3796, 35.6849}, {51.3923, 35.6882}, {51.3840, 35.6919}},
		Profile:     ProfileCar,
	})

	assert.Error(t, err)
	assert.Nil(t, res)
}
