package gosrm

import (
	"context"
)

// nearestServiceURL is the base path of OSRM nearest service.
const nearestServiceURL string = "/nearest/v1"

// NearestResponse is the response of OSRM's nearest service.
type NearestResponse struct {
	Response

	// Waypoints is an array of waypoint objects sorted by distance to the input coordinate.
	Waypoints []NearestWaypoint `json:"waypoints"`
}

// Nearset snaps a coordinate to the street network and returns the nearest n matches.
func Nearest(ctx context.Context, osrm OSRMClient, req Request, opts ...Option) (*NearestResponse, error) {
	u := req.buildURLPath(*osrm.baseURL, nearestServiceURL)

	osrm.applyOpts(u, opts)

	var res NearestResponse
	if err := osrm.get(ctx, u.String(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
