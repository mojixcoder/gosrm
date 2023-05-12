package gosrm

import (
	"fmt"
	"net/url"
)

type (
	// Option is the interface for adding options to requests.
	Option interface {
		apply(*url.URL)
	}

	// optionImpl is the type that implements Option interface.
	optionImpl func(*url.URL)
)

// apply implements the Option interface.
func (f optionImpl) apply(u *url.URL) {
	f(u)
}

// setQueryParam sets a query parameter in the URL.
func setQueryParam(u *url.URL, k, v string) {
	q := u.Query()
	q.Set(k, v)
	u.RawQuery = q.Encode()
}

// WithNumber sets number of nearest segments that should be returned.
// It should be >= 1.
// Can only used with nearest service.
func WithNumber(number uint8) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "number", fmt.Sprintf("%d", number))
	})
}

// WithAlternatives makes OSRM to search for alternative routes and return as well.
// Can be used in route service.
func WithAlternatives(alternatives bool) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "alternatives", fmt.Sprintf("%t", alternatives))
	})
}

// WithSteps makes OSRM to return route steps for each route leg.
// Can be used in route, match and trip services.
func WithSteps(steps bool) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "steps", fmt.Sprintf("%t", steps))
	})
}

// WithAnnotations makes OSRM to return additional metadata for each coordinate along the route geometry.
// Can be used in route, match and trip services.
func WithAnnotations(annotations Annotations) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "annotations", string(annotations))
	})
}

// WithGeometries sets the returned route geometry format (influences overview and per step).
// Can be used in route, match and trip services.
func WithGeometries(geometry Geometry) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "geometries", string(geometry))
	})
}

// WithOverview adds overview geometry either full, simplified according to highest zoom level it could be display on, or not at all.
// Can be used in route, match and trip services.
func WithOverview(overview Overview) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "overview", string(overview))
	})
}

// WithContinueStraight forces the route to keep going straight at waypoints constraining uturns there even if it would be faster.
// Default value depends on the profile.
// Can be used in route service.
func WithContinueStraight(cs ContinueStraight) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "continue_straight", string(cs))
	})
}

// WithSources uses location with given index as source.
// If the slice is empty uses all.
// Can be used in table service.
func WithSources(sources []uint16) Option {
	return optionImpl(func(u *url.URL) {
		if len(sources) == 0 {
			setQueryParam(u, "sources", "all")
		} else {
			setQueryParam(u, "sources", convertSliceToStr(sources, ";"))
		}
	})
}

// WithDestinations	uses location with given index as destination.
// If the slice is empty uses all.
// Can be used in table service.
func WithDestinations(destinations []uint16) Option {
	return optionImpl(func(u *url.URL) {
		if len(destinations) == 0 {
			setQueryParam(u, "destinations", "all")
		} else {
			setQueryParam(u, "destinations", convertSliceToStr(destinations, ";"))
		}
	})
}

// WithFallbackSpeed is used if no route found between a source/destination pair,
// calculate the as-the-crow-flies distance, then use this speed to estimate duration.
// should be greater than 0.
// Can be used in table service.
func WithFallbackSpeed(speed float64) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "fallback_speed", fmt.Sprintf("%f", speed))
	})
}

// WithFallbackCoordinate when using a fallback_speed,
// use the user-supplied coordinate (input), or the snapped location (snapped) for calculating distances.
// Can be used in table service.
func WithFallbackCoordinate(fc FallbackCoordinate) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "fallback_coordinate", string(fc))
	})
}

// WithScaleFactor should be uses in conjunction with annotations=durations. Scales the table duration values by this number.
// Can be used in table service.
func WithScaleFactor(sf float64) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "scale_factor", fmt.Sprintf("%f", sf))
	})
}

// WithTimestamps adds timestamps of the input locations in UNIX seconds.
// Timestamps need to be monotonically increasing.
// Can be used in match service.
func WithTimestamps(timestamps []int64) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "timestamps", convertSliceToStr(timestamps, ";"))
	})
}

// WithGaps allows the input track splitting based on huge timestamp gaps between points.
// Can be used in match service.
func WithGaps(gaps Gaps) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "gaps", string(gaps))
	})
}

// WithTidy allows the input track modification to obtain better matching quality for noisy tracks.
// Can be used in match service.
func WithTidy(tidy bool) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "tidy", fmt.Sprintf("%t", tidy))
	})
}

// WithWaypoints treats input coordinates indicated by given indices as waypoints in returned Match object.
// Default is to treat all input coordinates as waypoints.
// Can be used in route and match services.
func WithWaypoints(waypoints []uint16) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "waypoints", convertSliceToStr(waypoints, ";"))
	})
}

// WithRadiuses limits the search to given radius in meters.
// It's a general option and can be used in all services.
func WithRadiuses(radiuses []float32) Option {
	return optionImpl(func(u *url.URL) {
		if len(radiuses) == 0 {
			setQueryParam(u, "radiuses", "unlimited")
		} else {
			setQueryParam(u, "radiuses", convertSliceToStr(radiuses, ";"))
		}
	})
}

// WithRoundTrip is used when the returned route is a roundtrip (route returns to first location).
// Can be used in trip service.
func WithRoundTrip(isRound bool) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "roundtrip", fmt.Sprintf("%t", isRound))
	})
}

// WithSource is used when the returned route starts at any or first coordinate.
// Can be used in trip service.
func WithSource(source Source) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "source", string(source))
	})
}

// WithDestination is used when the returned route ends at any or last coordinate.
// Can be used in trip service.
func WithDestination(dest Destination) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "destination", string(dest))
	})
}

// WithCustomOption sets a custom option.
// Can be used if an option is not provided by the package.
func WithCustomOption(option, value string) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, option, value)
	})
}

// WithBearings limits the search to segments with given bearing in degrees towards true north in a clockwise direction.
// It's a general option and can be used in all services.
func WithBearings(bearings []Bearing) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "bearings", convertSliceToStr(bearings, ";"))
	})
}

// WithGenerateHints adds a hint to the response which can be used in subsequent requests.
// It's a general option and can be used in all services.
func WithGenerateHints(generate bool) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "generate_hints", fmt.Sprintf("%t", generate))
	})
}

// WithHints is hint from previous request to derive position in street network.
// Hint is a base64 string.
// It's a general option and can be used in all services.
func WithHints(hints []string) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "hints", convertSliceToStr(hints, ";"))
	})
}

// WithApproaches keeps waypoints on curbside.
// It's a general option and can be used in all services.
func WithApproaches(approaches []Approaches) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "approaches", convertSliceToStr(approaches, ";"))
	})
}

// WithExclude is an additive list of classes to avoid, the order does not matter.
// A class name determined by the profile or none.
// It's a general option and can be used in all services.
func WithExclude(classes []string) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "exclude", convertSliceToStr(classes, ";"))
	})
}

// WithSnapping default snapping avoids is_startpoint (see profile) edges, any will snap to any edge in the graph.
// It's a general option and can be used in all services.
func WithSnapping(snapping Snapping) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "snapping", string(snapping))
	})
}

// WithSkipWaypoints removes waypoints from the response.
// Waypoints are still calculated, but not serialized.
// Could be useful in case you are interested in some other part of the response and do not want to transfer waste data.
// It's a general option and can be used in all services.
func WithSkipWaypoints(skip bool) Option {
	return optionImpl(func(u *url.URL) {
		setQueryParam(u, "skip_waypoints", fmt.Sprintf("%t", skip))
	})
}
