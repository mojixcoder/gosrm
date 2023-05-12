package gosrm

import (
	"bytes"
	"fmt"
	"strings"
)

// convertSliceToStr converts a slice of anything to {v}{sep}{v} representation of the slice.
func convertSliceToStr(s any, sep string) string {
	var b bytes.Buffer

	switch slice := s.(type) {
	case []uint16:
		for _, n := range slice {
			b.WriteString(fmt.Sprintf("%d%s", n, sep))
		}
	case []int64:
		for _, n := range slice {
			b.WriteString(fmt.Sprintf("%d%s", n, sep))
		}
	case []float32:
		for _, f := range slice {
			b.WriteString(fmt.Sprintf("%f%s", f, sep))
		}
	case []string:
		for _, s := range slice {
			b.WriteString(s + sep)
		}
	case []Approaches:
		for _, approach := range slice {
			b.WriteString(string(approach) + sep)
		}
	case []Bearing:
		for _, bearing := range slice {
			b.WriteString(fmt.Sprintf("%d,%d%s", bearing.Value, bearing.Range, sep))
		}
	case []Coordinate:
		for _, coor := range slice {
			b.WriteString(fmt.Sprintf("%f,%f%s", coor[0], coor[1], sep))
		}
	default:
		panic("unsupported type")
	}

	return strings.TrimSuffix(b.String(), sep)
}
