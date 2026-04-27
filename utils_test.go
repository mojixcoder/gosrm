package gosrm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertSliceToStr(t *testing.T) {
	t.Run("[]uint16", func(t *testing.T) {
		assert.Equal(t, "1;2;3", convertSliceToStr([]uint16{1, 2, 3}, ";"))
	})
	t.Run("[]int64", func(t *testing.T) {
		assert.Equal(t, "1;2;3", convertSliceToStr([]int64{1, 2, 3}, ";"))
	})
	t.Run("[]float32", func(t *testing.T) {
		assert.Equal(t, "1.000000;2.000000;3.000000", convertSliceToStr([]float32{1, 2, 3}, ";"))
	})
	t.Run("[]string", func(t *testing.T) {
		assert.Equal(t, "abc;qwe", convertSliceToStr([]string{"abc", "qwe"}, ";"))
	})
	t.Run("[]Approaches", func(t *testing.T) {
		assert.Equal(t, "curb", convertSliceToStr([]Approaches{ApproachesCurb}, ";"))
	})
	t.Run("[]Bearing", func(t *testing.T) {
		assert.Equal(t, "150,100;200,100", convertSliceToStr([]Bearing{{Value: 150, Range: 100}, {Value: 200, Range: 100}}, ";"))
	})
	t.Run("[]Coordinate", func(t *testing.T) {
		assert.Equal(t, "1.000000,2.000000;10.000000,20.000000", convertSliceToStr([]Coordinate{{1, 2}, {10, 20}}, ";"))
	})
}
