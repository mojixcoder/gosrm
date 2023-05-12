package gosrm

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	osrm, err := New(getOSRMAddress())
	assert.NoError(t, err)

	res, err := Route[string](context.Background(), osrm, Request{
		Coordinates: []Coordinate{{49.5270, 31.4505}, {49.5277, 31.4521}},
		Profile:     ProfileCar,
	})

	assert.NoError(t, err)
	assert.Equal(t, CodeOK, res.Code)
	assert.Len(t, res.Routes, 1)

	osrm.baseURL.Host = "invalid"
	res, err = Route[string](context.Background(), osrm, Request{
		Coordinates: []Coordinate{{49.5270, 31.4505}, {49.5277, 31.4521}},
		Profile:     ProfileCar,
	})

	assert.Error(t, err)
	assert.Nil(t, res)
}
