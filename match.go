package gosrm

import (
	"context"
)

// matchServiceURL is the base path of OSRM match service.
const matchServiceURL string = "/match/v1"

// MatchResponse is the response of OSRM's match service.
type MatchResponse[T GeometryType] struct {
	Response

	// Tracepoints is an array of waypoint objects representing all points of the trace in order.
	// If the trace point was ommited by map matching because it is an outlier, the entry will be null.
	Tracepoints []Tracepoint `json:"tracepoints"`

	// Matchings is an array of route objects that assemble the trace.
	Matchings []Matching[T] `json:"matchings"`
}

// Match matches/snaps given GPS points to the road network in the most plausible way.
func Match[T GeometryType](ctx context.Context, osrm OSRMClient, req Request, opts ...Option) (*MatchResponse[T], error) {
	u := req.buildURLPath(*osrm.baseURL, matchServiceURL)

	osrm.applyOpts(u, opts)

	var res MatchResponse[T]
	if err := osrm.get(ctx, u.String(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
