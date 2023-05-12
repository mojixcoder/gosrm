package gosrm

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionImpl_apply(t *testing.T) {
	var u url.URL
	f := optionImpl(func(u *url.URL) {
		u.Path = "/test"
	})

	f.apply(&u)

	assert.Equal(t, "/test", u.Path)
}

func TestSetQueryParam(t *testing.T) {
	var u url.URL

	setQueryParam(&u, "key", "value")

	assert.Equal(t, "key=value", u.RawQuery)
	assert.Equal(t, "value", u.Query().Get("key"))
}

func TestOptions(t *testing.T) {
	var osrm OSRMClient
	var u url.URL

	opts := []Option{
		WithNumber(3),
		WithAlternatives(true),
		WithSteps(false),
		WithAnnotations(AnnotationsSpeed),
		WithGeometries(GeometryGeoJSON),
		WithOverview(OverviewSimplified),
		WithContinueStraight(ContinueStraightTrue),
		WithSources([]uint16{0, 1, 2}),
		WithDestinations([]uint16{1}),
		WithTimestamps([]int64{1234, 5721}),
		WithGaps(GapsIgnore),
		WithTidy(true),
		WithWaypoints([]uint16{0, 1}),
		WithRadiuses([]float32{1.567, 2.5683}),
		WithRoundTrip(false),
		WithSource(SourceAny),
		WithDestination(DestinationLast),
		WithBearings([]Bearing{{Value: 200, Range: 50}}),
		WithGenerateHints(true),
		WithHints([]string{"asd", "qwe"}),
		WithApproaches([]Approaches{ApproachesUnrestricted}),
		WithExclude([]string{"motorway"}),
		WithSnapping(SnappingDefault),
		WithCustomOption("opt", "val"),
		WithSkipWaypoints(false),
		WithFallbackSpeed(1.432123),
		WithScaleFactor(1),
		WithFallbackCoordinate(FallbackCoordinateInput),
	}
	osrm.applyOpts(&u, opts)
	q := u.Query()

	assert.Equal(t, "3", q.Get("number"))
	assert.Equal(t, "true", q.Get("alternatives"))
	assert.Equal(t, "false", q.Get("steps"))
	assert.Equal(t, string(AnnotationsSpeed), q.Get("annotations"))
	assert.Equal(t, string(GeometryGeoJSON), q.Get("geometries"))
	assert.Equal(t, string(OverviewSimplified), q.Get("overview"))
	assert.Equal(t, string(ContinueStraightTrue), q.Get("continue_straight"))
	assert.Equal(t, "0;1;2", q.Get("sources"))
	assert.Equal(t, "1", q.Get("destinations"))
	assert.Equal(t, "1234;5721", q.Get("timestamps"))
	assert.Equal(t, "ignore", q.Get("gaps"))
	assert.Equal(t, "true", q.Get("tidy"))
	assert.Equal(t, "0;1", q.Get("waypoints"))
	assert.Equal(t, "1.567000;2.568300", q.Get("radiuses"))
	assert.Equal(t, "false", q.Get("roundtrip"))
	assert.Equal(t, "any", q.Get("source"))
	assert.Equal(t, "last", q.Get("destination"))
	assert.Equal(t, "200,50", q.Get("bearings"))
	assert.Equal(t, "true", q.Get("generate_hints"))
	assert.Equal(t, "asd;qwe", q.Get("hints"))
	assert.Equal(t, "unrestricted", q.Get("approaches"))
	assert.Equal(t, "motorway", q.Get("exclude"))
	assert.Equal(t, "default", q.Get("snapping"))
	assert.Equal(t, "val", q.Get("opt"))
	assert.Equal(t, "false", q.Get("skip_waypoints"))
	assert.Equal(t, "1.432123", q.Get("fallback_speed"))
	assert.Equal(t, "1.000000", q.Get("scale_factor"))
	assert.Equal(t, string(FallbackCoordinateInput), q.Get("fallback_coordinate"))

	opts = []Option{
		WithSources(nil),
		WithDestinations(nil),
		WithRadiuses(nil),
	}
	osrm.applyOpts(&u, opts)
	q = u.Query()

	assert.Equal(t, "all", q.Get("sources"))
	assert.Equal(t, "all", q.Get("destinations"))
	assert.Equal(t, "unlimited", q.Get("radiuses"))
}
