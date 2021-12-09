package _9

import (
	"2021/common/reader"
	"github.com/stretchr/testify/assert"
	"testing"
)

var expectedLowPoints = []Coordinates{{0, 1}, {0, 9}, {2, 2}, {4, 6}}

func TestHeightMap_LowPoints(t *testing.T) {
	h, err := reader.Read2DIntSliceFile("testdata/09.log")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, expectedLowPoints, HeightMap(h).LowPoints())
}

func TestEvaluateRisks(t *testing.T) {
	h, err := reader.Read2DIntSliceFile("testdata/09.log")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 15, EvaluateRisks(h, expectedLowPoints))
}

func TestCoordinates_FindBasin(t *testing.T) {
	d, err := reader.Read2DIntSliceFile("testdata/09.log")
	if err != nil {
		t.Fail()
	}

	h := HeightMap(d)
	b := h.Basins()

	assert.Equal(t, 4, len(b))
	assert.Equal(t, 3, b[0].Size())
	assert.Equal(t, 9, b[1].Size())
	assert.Equal(t, 14, b[2].Size())
	assert.Equal(t, 9, b[3].Size())
}

func TestMultiplyBasinsSizes(t *testing.T) {
	d, err := reader.Read2DIntSliceFile("testdata/09.log")
	if err != nil {
		t.Fail()
	}

	h := HeightMap(d)
	b := GetTop3LargestBasins(h.Basins())

	assert.Equal(t, 1134, MultiplyBasinsSizes(b))
}
