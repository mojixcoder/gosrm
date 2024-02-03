package gosrm

import "time"

type (
	// Coordinate is a {Lng, Lat} point.
	Coordinate [2]float64

	// Response is the common fields which are used in all of the OSRM services responses.
	Response struct {
		// Code is the error code returned by OSRM.
		Code Code `json:"code"`

		// Message is a optional human-readable error message.
		Message string `json:"message"`

		DataVersion time.Time `json:"data_version"`
	}

	// GeometryType is the type used to represent multiple geometry formats.
	// polyline5, polyline6, geojson.
	GeometryType interface {
		string | LineString
	}

	// Bearing is the {value},{range} bearing type.
	Bearing struct {
		// Value is an integer 0 to 360.
		Value uint16

		// Range is an integer 0 to 180.
		Range uint16
	}

	// LineString is a GeoJSON line string.
	LineString struct {
		// Type of the line string.
		Type string `json:"type"`

		// Coordinates of the line string.
		Coordinates []Coordinate `json:"coordinates"`
	}

	// Waypoint is the object used to describe waypoint on a route.
	Waypoint struct {
		// Name of the street the coordinate snapped to
		Name string `json:"name"`

		// Hint is an unique internal identifier of the segment (ephemeral, not constant over data updates)
		// This can be used on subsequent request to significantly speed up the query and to connect multiple services.
		// E.g. you can use the hint value obtained by the nearest query as hint values for route inputs.
		Hint string `json:"hint"`

		// Distance of the snapped point from the original, in meters.
		Distance float32 `json:"distance"`

		// Location is an array that contains the [longitude, latitude] pair of the snapped coordinate
		Location Coordinate `json:"location"`
	}

	// NearestWaypoint is the object used to describe nearest waypoint on a route.
	NearestWaypoint struct {
		Waypoint

		// Nodes is an array of OpenStreetMap node ids.
		Nodes []uint64 `json:"nodes"`
	}

	// TripWaypoint is the object used to describe trip waypoint on a route.
	TripWaypoint struct {
		Waypoint

		// TripsIndex is the index to trips of the sub-trip the point was matched to.
		TripsIndex uint16 `json:"trips_index"`

		// WaypointIndex is the index of the point in the trip.
		WaypointIndex uint16 `json:"waypoint_index"`
	}

	// Lane represents a turn lane at the corresponding turn location.
	Lane struct {
		// Indications are the indications (e.g. marking on the road) specifying the turn lane.
		Indications []string `json:"indications"`

		// Valid is a boolean flag indicating whether the lane is a valid choice in the current maneuver.
		Valid bool `json:"valid"`
	}

	// Intersection gives a full representation of any cross-way the path passes bay.
	// For every step, the very first intersection (intersections[0]) corresponds to the location of the StepManeuver.
	// Further intersections are listed for every cross-way until the next turn instruction.
	Intersection struct {
		// Location is a [longitude, latitude] pair describing the location of the turn.
		Location Coordinate `json:"location"`

		// Bearings is a list of bearing values (e.g. [0,90,180,270]) that are available at the intersection.
		// The bearings describe all available roads at the intersection.
		Bearings []uint16 `json:"bearings"`

		// Classes is an array of strings signifying the classes (as specified in the profile) of the road exiting the intersection.
		Classes []string `json:"classes"`

		// Entry is a list of entry flags, corresponding in a 1:1 relationship to the bearings.
		// A value of true indicates that the respective road could be entered on a valid route.
		// false indicates that the turn onto the respective road would violate a restriction.
		Entry []bool `json:"entry"`

		// In is the index 	into bearings/entry array. Used to calculate the bearing just before the turn.
		In uint16 `json:"in"`

		// Out is the index into the bearings/entry array. Used to extract the bearing just after the turn.
		Out uint16 `json:"out"`

		// Lanes is an array of Lane objects that denote the available turn lanes at the intersection.
		Lanes []Lane `json:"lanes"`
	}

	// StepManeuver holds information about maneuver in a step.
	StepManeuver struct {
		// Location is a [longitude, latitude] pair describing the location of the turn.
		Location Coordinate `json:"location"`

		// BearingBefore is the clockwise angle from true north to the direction of travel immediately before the maneuver.
		BearingBefore float32 `json:"bearing_before"`

		// BearingAfter is the clockwise angle from true north to the direction of travel immediately before the maneuver.
		BearingAfter float32 `json:"bearing_after"`

		// Type is a string indicating the type of maneuver.
		Type string `json:"type"`

		// Modifier is an optional string indicating the direction change of the maneuver.
		Modifier string `json:"modifier"`

		// Exit is an optional integer indicating number of the exit to take.
		Exit uint16 `json:"exit"`
	}

	// RouteStep consists of a maneuver such as a turn or merge,
	// followed by a distance of travel along a single way to the subsequent step.
	RouteStep[T GeometryType] struct {
		// Distance is the distance of travel from the maneuver to the subsequent step, in meters.
		Distance float32 `json:"distance"`

		// Duration is the estimated travel time, in float number of seconds.
		Duration float32 `json:"duration"`

		// Weight is the calculated weight of the step.
		Weight float32 `json:"weight"`

		// Exits is the exit numbers or names of the way. Will be undefined if there are no exit numbers or names.
		Exits uint16 `json:"exits"`

		// Name of the way along which travel proceeds.
		Name string `json:"name"`

		// Ref a reference number or code for the way. Optionally included, if ref data is available for the given way.
		Ref string `json:"ref"`

		// Pronunciation is the pronunciation hint of the way name. Will be undefined if there is no pronunciation hit.
		Pronunciation string `json:"pronunciation"`

		// RotaryName is the name for the rotary. Optionally included, if the step is a rotary and a rotary name is available.
		RotaryName string `json:"rotary_name"`

		// RotaryPronunciation is the pronunciation hint of the rotary name.
		// Optionally included, if the step is a rotary and a rotary pronunciation is available.
		RotaryPronunciation string `json:"rotary_pronunciation"`

		// DrivingSide is the legal driving side at the location for this step. Either left or right.
		DrivingSide string `json:"driving_side"`

		// Destinations of the way. Will be undefined if there are no destinations.
		Destinations []Waypoint `json:"destinations"`

		// Mode is a string signifying the mode of transportation.
		Mode string `json:"mode"`

		// Intersections is a list of Intersection objects that are passed along the segment, the very first belonging to the StepManeuver.
		Intersections []Intersection `json:"intersections"`

		// Maneuver represents the maneuver.
		Maneuver StepManeuver `json:"maneuver"`

		// Geometry is the unsimplified geometry of the route segment, depending on the geometries parameter.
		// if geometries is geojson	use LineString otherwise use string.
		Geometry T `json:"geometry"`
	}

	// Metadata related to other annotations.
	Metadata struct {
		// DataSourcesNames is the names of the data sources used for the speed between each pair of coordinates.
		DataSourcesNames []string `json:"datasource_names"`
	}

	// Annotation of the whole route leg with fine-grained information about each segment or node id.
	Annotation struct {
		// Distance is the distance, in metres, between each pair of coordinates.
		Distance []float32 `json:"distance"`

		// Duration is the duration between each pair of coordinates, in seconds.
		Duration []float32 `json:"duration"`

		// DataSources is the index of the datasource for the speed between each pair of coordinates.
		// 0 is the default profile, other values are supplied via --segment-speed-file to osrm-contract.
		DataSources []uint16 `json:"datasources"`

		// Nodes is the OSM node ID for each coordinate along the route, excluding the first/last user-supplied coordinates.
		Nodes []uint64 `json:"nodes"`

		// Weight is the weights between each pair of coordinates. Does not include any turn costs.
		Weight []float32 `json:"Weight"`

		// Speed is the convenience field, calculation of distance / duration rounded to one decimal place.
		Speed []float32 `json:"speed"`

		// Metadata related to other annotations
		Metadata Metadata `json:"metadata"`
	}

	// RouteLeg represents a route between two waypoints.
	RouteLeg[T GeometryType] struct {
		// Distance is the distance traveled by this route leg, in meters.
		Distance float32 `json:"distance"`

		// Duration is the estimated travel time, in seconds.
		Duration float32 `json:"duration"`

		// Summary of the route taken as string. Depends on the steps parameter.
		Summary string `json:"summary"`

		// Weight is the calculated weight of the route leg.
		Weight float32 `json:"weight"`

		// Annotation is additional details about each coordinate along the route geometry.
		Annotation Annotation `json:"annotation"`

		// Steps are route steps of the route leg.
		// if geometries is geojson	use LineString otherwise use string.
		Steps []RouteStep[T] `json:"steps"`
	}

	// RouteType represents a route through (potentially multiple) waypoints.
	RouteType[T GeometryType] struct {
		// Distance is the distance traveled by the route, in meters.
		Distance float32 `json:"distance"`

		// Duration	is the estimated travel time, in seconds.
		Duration float32 `json:"duration"`

		// Weight is the calculated weight of the route.
		Weight float32 `json:"weight"`

		// WeightName is the name of the weight profile used during the extraction phase.
		WeightName string `json:"weight_name"`

		// Legs are the legs between the given waypoints.
		// if geometries is geojson	use LineString otherwise use string.
		Legs []RouteLeg[T] `json:"legs"`

		// Geometry is the whole geometry of the route value depending on overview parameter, format depending on the geometries parameter.
		// if geometries is geojson	use LineString otherwise use string.
		Geometry T `json:"geometry"`
	}

	// Tracepoint is a waypoint object representing a point of the trace.
	Tracepoint struct {
		Waypoint

		// WaypointIndex is the index of the waypoint inside the matched route.
		WaypointIndex uint16 `json:"waypoint_index"`

		// MatchingIndex is the index to the Route object in matchings the sub-trace was matched to.
		MatchingIndex uint16 `json:"matchings_index"`

		// AlternativesCount is the number of probable alternative matchings for this tracepoint.
		// A value of zero indicates that this point was matched unambiguously.
		// Split the trace at these points for incremental map matching.
		AlternativesCount uint16 `json:"alternatives_count"`
	}

	// Matching is a route object that assembles the trace.
	Matching[T GeometryType] struct {
		RouteType[T]

		// Confidence of the matching. float value between 0 and 1. 1 is very confident that the matching is correct.
		Confidence float32 `json:"confidence"`
	}
)

// IsOk returns true if request could be processed as expected by OSRM.
func (res Response) IsOk() bool {
	return res.Code == CodeOK
}
