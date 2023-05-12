package gosrm

import (
	"context"
)

// tripServiceURL is the base path of OSRM trip service.
const tripServiceURL string = "/trip/v1"

// TripResponse is the response of OSRM's trip service.
type TripResponse[T GeometryType] struct {
	Response

	// Waypoints is an array of waypoint objects representing all waypoints in input order.
	Waypoints []TripWaypoint `json:"waypoints"`

	// Trips is an array of Route objects that assemble the trace.
	Trips []RouteType[T] `json:"trips"`
}

// Trip service solves the Traveling Salesman Problem using a greedy heuristic (farthest-insertion algorithm)
// for 10 or more waypoints and uses brute force for less than 10 waypoints.
// The returned path does not have to be the fastest one. As TSP is NP-hard it only returns an approximation.
// Note that all input coordinates have to be connected for the trip service to work.
func Trip[T GeometryType](ctx context.Context, osrm OSRMClient, req Request, opts ...Option) (*TripResponse[T], error) {
	u := req.buildURLPath(*osrm.baseURL, tripServiceURL)

	osrm.applyOpts(u, opts)

	var res TripResponse[T]
	if err := osrm.get(ctx, u.String(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
