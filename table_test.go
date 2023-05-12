package gosrm

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	osrm, err := New(getOSRMAddress())
	assert.NoError(t, err)

	res, err := Table(context.Background(), osrm, Request{
		Coordinates: []Coordinate{{51.6488, 32.6905}, {51.6450, 32.6920}},
		Profile:     ProfileCar,
	})

	assert.NoError(t, err)
	assert.Equal(t, CodeOK, res.Code)
	assert.Len(t, res.Durations, 2)
	assert.Len(t, res.Destinations, 2)
	assert.Len(t, res.Sources, 2)

	osrm.baseURL.Host = "invalid"
	res, err = Table(context.Background(), osrm, Request{
		Coordinates: []Coordinate{{51.6488, 32.6905}, {51.6450, 32.6920}},
		Profile:     ProfileCar,
	})

	assert.Error(t, err)
	assert.Nil(t, res)
}
