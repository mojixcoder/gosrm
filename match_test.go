package gosrm

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	osrm, err := New(getOSRMAddress())
	assert.NoError(t, err)

	res, err := Match[string](context.Background(), osrm, Request{
		Coordinates: []Coordinate{{51.2831, 35.7676}, {51.2811, 35.7665}},
		Profile:     ProfileCar,
	})

	assert.NoError(t, err)
	assert.Equal(t, CodeOK, res.Code)
	assert.Len(t, res.Tracepoints, 2)

	osrm.baseURL.Host = "invalid"
	res, err = Match[string](context.Background(), osrm, Request{
		Coordinates: []Coordinate{{51.2831, 35.7676}, {51.2811, 35.7665}},
		Profile:     ProfileCar,
	})

	assert.Error(t, err)
	assert.Nil(t, res)
}
