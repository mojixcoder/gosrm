package gosrm

import (
	"bytes"
	"fmt"
	"strings"
)

// sliceElement is the type constraint for elements supported by convertSliceToStr.
type sliceElement interface {
	uint16 | int64 | float32 | string | Approaches | Bearing | Coordinate
}

// convertSliceToStr converts a slice to a {v}{sep}{v} string representation.
func convertSliceToStr[T sliceElement](s []T, sep string) string {
	var b bytes.Buffer

	for _, v := range s {
		switch val := any(v).(type) {
		case uint16:
			b.WriteString(fmt.Sprintf("%d%s", val, sep))
		case int64:
			b.WriteString(fmt.Sprintf("%d%s", val, sep))
		case float32:
			b.WriteString(fmt.Sprintf("%f%s", val, sep))
		case string:
			b.WriteString(val + sep)
		case Approaches:
			b.WriteString(string(val) + sep)
		case Bearing:
			b.WriteString(fmt.Sprintf("%d,%d%s", val.Value, val.Range, sep))
		case Coordinate:
			b.WriteString(fmt.Sprintf("%f,%f%s", val[0], val[1], sep))
		}
	}

	return strings.TrimSuffix(b.String(), sep)
}
