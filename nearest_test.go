package gosrm

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNearest(t *testing.T) {
	osrm, err := New(getOSRMAddress())
	assert.NoError(t, err)

	res, err := Nearest(context.Background(), osrm, Request{
		Coordinates: []Coordinate{{49.5257, 31.4507}},
		Profile:     ProfileCar,
	}, WithNumber(2))

	assert.NoError(t, err)
	assert.Equal(t, CodeOK, res.Code)
	assert.Len(t, res.Waypoints, 2)

	osrm.baseURL.Host = "invalid"
	res, err = Nearest(context.Background(), osrm, Request{
		Coordinates: []Coordinate{{49.5257, 31.4507}},
		Profile:     ProfileCar,
	}, WithNumber(2))

	assert.Error(t, err)
	assert.Nil(t, res)
}
