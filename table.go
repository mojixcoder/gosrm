package gosrm

import (
	"context"
)

// tableServiceURL is the base path of OSRM route service.
const tableServiceURL string = "/table/v1"

// TableResponse is the response of OSRM's table service.
type TableResponse struct {
	Response

	// Durations is an array of arrays that stores the matrix in row-major order.
	// durations[i][j] gives the travel time from the i-th waypoint to the j-th waypoint, in seconds.
	Durations [][]float32 `json:"durations"`

	// Distances is an array of arrays that stores the matrix in row-major order.
	// distances[i][j] gives the travel distance from the i-th source to the j-th destination, in meters.
	Distances [][]float32 `json:"distances"`

	// Destinations is an array of Waypoint objects describing all destinations in order.
	Destinations []Waypoint `json:"destinations"`

	// Sources is an array of Waypoint objects describing all sources in order.
	Sources []Waypoint `json:"sources"`

	// FallbackSpeedCells is an optional array of arrays containing i,j pairs indicating
	// which cells contain estimated values based on fallback_speed.
	// Will be absent if fallback_speed is not used.
	FallbackSpeedCells [][]uint16 `json:"fallback_speed_cells"`
}

// Table computes the duration of the fastest route between all pairs of supplied coordinates.
func Table(ctx context.Context, osrm OSRMClient, req Request, opts ...Option) (*TableResponse, error) {
	u := req.buildURLPath(*osrm.baseURL, tableServiceURL)

	osrm.applyOpts(u, opts)

	var res TableResponse
	if err := osrm.get(ctx, u.String(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
