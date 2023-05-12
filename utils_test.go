package gosrm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertSliceToStr(t *testing.T) {
	assert.PanicsWithValue(t, "unsupported type", func() {
		convertSliceToStr(nil, "")
	})

	testCases := []struct {
		expectedRes, name, sep string
		input                  any
	}{
		{name: "[]uint16", expectedRes: "1;2;3", sep: ";", input: []uint16{1, 2, 3}},
		{name: "[]int64", expectedRes: "1;2;3", sep: ";", input: []int64{1, 2, 3}},
		{name: "[]float32", expectedRes: "1.000000;2.000000;3.000000", sep: ";", input: []float32{1, 2, 3}},
		{name: "[]string", expectedRes: "abc;qwe", sep: ";", input: []string{"abc", "qwe"}},
		{name: "[]Approaches", expectedRes: "curb", sep: ";", input: []Approaches{ApproachesCurb}},
		{name: "[]Bearing", expectedRes: "150,100;200,100", sep: ";", input: []Bearing{{Value: 150, Range: 100}, {Value: 200, Range: 100}}},
		{name: "[]Coordinate", expectedRes: "1.000000,2.000000;10.000000,20.000000", sep: ";", input: []Coordinate{{1, 2}, {10, 20}}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := convertSliceToStr(testCase.input, testCase.sep)

			assert.Equal(t, testCase.expectedRes, res)
		})
	}
}
