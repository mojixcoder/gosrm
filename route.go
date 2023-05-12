package gosrm

import (
	"context"
)

// routeServiceURL is the base path of OSRM route service.
const routeServiceURL string = "/route/v1"

// RouteResponse is the response of OSRM's route service.
type RouteResponse[T GeometryType] struct {
	Response

	Routes    []RouteType[T] `json:"routes"`
	Waypoints []Waypoint     `json:"waypoints"`
}

// Route finds the fastest route between coordinates in the supplied order.
func Route[T GeometryType](ctx context.Context, osrm OSRMClient, req Request, opts ...Option) (*RouteResponse[T], error) {
	u := req.buildURLPath(*osrm.baseURL, routeServiceURL)

	osrm.applyOpts(u, opts)

	var res RouteResponse[T]
	if err := osrm.get(ctx, u.String(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
