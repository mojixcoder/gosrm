package gosrm

type (
	// Geometry is the geometry format.
	Geometry string

	// Overview is the type for overview option.
	Overview string

	// ContinueStraight is the type for continue straight option.
	ContinueStraight string

	// Gaps is the type for gaps option.
	Gaps string

	// Annotations is the type for annotations option.
	Annotations string

	// Source is the type for source option.
	Source string

	// Destination is the type for destination option.
	Destination string

	// Code is the error code returned by OSRM.
	Code string

	// Profile is the mode of transportation, is determined statically by the Lua profile that is used to prepare the data using osrm-extract.
	// Typically car, bike or foot if using one of the supplied profiles.
	Profile string

	// Approaches is the type for approaches option.
	Approaches string

	// Snapping is the type for snapping option.
	Snapping string

	// FallbackCoordinate is the type for fallback coordinate option.
	FallbackCoordinate string
)

const (
	// CodeOK is returned when request could be processed as expected.
	CodeOK Code = "Ok"

	//CodeInvalidUrl is returned when URL string is invalid.
	CodeInvalidUrl Code = "InvalidUrl"

	// CodeInvalidService is returned when service name is invalid.
	CodeInvalidService Code = "InvalidService"

	// CodeInvalidVersion is returned when version is not found.
	CodeInvalidVersion Code = "InvalidVersion"

	// CodeInvalidOptions is returned when options are invalid.
	CodeInvalidOptions Code = "InvalidOptions"

	// CodeInvalidQuery is returned when the query string is synctactically malformed.
	CodeInvalidQuery Code = "InvalidQuery"

	// CodeInvalidValue is returned when the successfully parsed query parameters are invalid.
	CodeInvalidValue Code = "InvalidValue"

	// CodeNoSegment is returned when one of the supplied input coordinates could not snap to street segment.
	CodeNoSegment Code = "NoSegment"

	// CodeTooBig is returned when the request size violates one of the service specific request size restrictions.
	CodeTooBig Code = "TooBig"

	// CodeNoRoute is returned when no route was found.
	CodeNoRoute Code = "NoRoute"

	// CodeNoTable is returned when no route was found.
	CodeNoTable Code = "NoTable"

	// CodeNoMatch is returned when no match was found.
	CodeNoMatch Code = "NoMatch"

	// CodeNoTrips is returned when no trips were found because input coordinates are not connected.
	CodeNoTrips Code = "NoTrips"

	// CodeNotImplemented is returned when this request is not supported.
	CodeNotImplemented Code = "NotImplemented"
)

const (
	// ProfileDriving is the driving transportation mode.
	ProfileDriving Profile = "driving"

	// ProfileCar is the car transportation mode.
	ProfileCar Profile = "car"

	// ProfileBike is the bike transportation mode.
	ProfileBike Profile = "bike"

	// ProfileFoot is the foot transportation mode.
	ProfileFoot Profile = "foot"
)

const (
	// GeometryPolyline is the geometry polyline type.
	GeometryPolyline Geometry = "polyline"

	// GeometryPolyline6 is the geometry polyline6 type.
	GeometryPolyline6 Geometry = "polyline6"

	// GeometryGeoJSON is the geometry geojson type.
	GeometryGeoJSON Geometry = "geojson"
)

const (
	// OverviewSimplified is the simplified overview.
	OverviewSimplified Overview = "simplified"

	// OverviewFull is the full overview.
	OverviewFull Overview = "full"

	// OverviewFalse disables the overview.
	OverviewFalse Overview = "false"
)

const (
	// ContinueStraightDefault is the default continue straight option value.
	ContinueStraightDefault ContinueStraight = "default"

	// ContinueStraightTrue enables continue straight.
	ContinueStraightTrue ContinueStraight = "true"

	// ContinueStraightFalse disables continue straight.
	ContinueStraightFalse ContinueStraight = "false"
)

const (
	// GapsSplit is the default gaps option value.
	GapsSplit Gaps = "split"

	// GapsIgnore is the ignore gaps option.
	GapsIgnore Gaps = "ignore"
)

const (
	// AnnotationsTrue enables annotations.
	AnnotationsTrue Annotations = "true"

	// AnnotationsFalse disables annotations.
	AnnotationsFalse Annotations = "false"

	// AnnotationsNodes is node annotations.
	AnnotationsNodes Annotations = "nodes"

	// AnnotationsSpeed is speed annotations.
	AnnotationsSpeed Annotations = "speed"

	// AnnotationsWeight is weight annotations.
	AnnotationsWeight Annotations = "weight"

	// AnnotationsDistance is distance annotations.
	AnnotationsDistance Annotations = "distance"

	// AnnotationsDuration is duration annotations.
	AnnotationsDuration Annotations = "duration"

	// AnnotationsDataSources is data sources annotations.
	AnnotationsDataSources Annotations = "datasources"

	// AnnotationsDurationDistance is duration and distance annotations.
	AnnotationsDurationDistance Annotations = "duration,distance"
)

const (
	// SourceAny is the any source option.
	SourceAny Source = "any"

	// SourceFirst is the first source option.
	SourceFirst Source = "first"
)

const (
	// DestinationAny is the any destination option.
	DestinationAny Destination = "any"

	// DestinationLast is the last destination option.
	DestinationLast Destination = "last"
)

const (
	// ApproachesCurb is the curb approaches.
	ApproachesCurb Approaches = "curb"

	// ApproachesUnrestricted is the unrestricted approaches.
	ApproachesUnrestricted Approaches = "unrestricted"
)

const (
	// SnappingDefault is the default snapping.
	SnappingDefault Snapping = "default"

	// SnappingAny is the any snapping.
	SnappingAny Snapping = "any"
)

const (
	// FallbackCoordinateInput when using a fallback_speed, use the user-supplied coordinate (input).
	FallbackCoordinateInput FallbackCoordinate = "input"

	// FallbackCoordinateSnapped when using a fallback_speed, use the snapped location (snapped) for calculating distances.
	FallbackCoordinateSnapped FallbackCoordinate = "snapped"
)
