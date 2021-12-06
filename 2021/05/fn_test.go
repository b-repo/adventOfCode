package _5

import (
	"2021/common/reader"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

var expectedLines = []Line{
	{Start: Coordinate{X: 0, Y: 9}, End: Coordinate{X: 5, Y: 9}},
	{Start: Coordinate{X: 8, Y: 0}, End: Coordinate{X: 0, Y: 8}},
	{Start: Coordinate{X: 9, Y: 4}, End: Coordinate{X: 3, Y: 4}},
	{Start: Coordinate{X: 2, Y: 2}, End: Coordinate{X: 2, Y: 1}},
	{Start: Coordinate{X: 7, Y: 0}, End: Coordinate{X: 7, Y: 4}},
	{Start: Coordinate{X: 6, Y: 4}, End: Coordinate{X: 2, Y: 0}},
	{Start: Coordinate{X: 0, Y: 9}, End: Coordinate{X: 2, Y: 9}},
	{Start: Coordinate{X: 3, Y: 4}, End: Coordinate{X: 1, Y: 4}},
	{Start: Coordinate{X: 0, Y: 0}, End: Coordinate{X: 8, Y: 8}},
	{Start: Coordinate{X: 5, Y: 5}, End: Coordinate{X: 8, Y: 2}},
}

var expectedFilteredLines = []Line{
	{Start: Coordinate{X: 0, Y: 9}, End: Coordinate{X: 5, Y: 9}},
	{Start: Coordinate{X: 9, Y: 4}, End: Coordinate{X: 3, Y: 4}},
	{Start: Coordinate{X: 2, Y: 2}, End: Coordinate{X: 2, Y: 1}},
	{Start: Coordinate{X: 7, Y: 0}, End: Coordinate{X: 7, Y: 4}},
	{Start: Coordinate{X: 0, Y: 9}, End: Coordinate{X: 2, Y: 9}},
	{Start: Coordinate{X: 3, Y: 4}, End: Coordinate{X: 1, Y: 4}},
}

func TestParseLines(t *testing.T) {
	data, err := reader.ReadStringFileAndSplitByLine("testdata/05.log")
	if err != nil {
		panic(err.Error())
	}

	l, m := ParseLines(data, func(l Line) bool {
		return l.Start.X == l.End.X || l.Start.Y == l.End.Y
	})

	assert.Equal(t, 9, m)

	assert.Equal(t, expectedFilteredLines, l)
}

func TestInitField(t *testing.T) {
	e := Field{
		Map: [][]int{
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		},
		Overlap:    map[int]int{},
		MaxOverlap: 0,
	}

	f := InitField(5)

	assert.Equal(t, &e, f)
}

func TestDrawLines(t *testing.T) {
	data, err := reader.ReadStringFileAndSplitByLine("testdata/05.log")
	if err != nil {
		panic(err.Error())
	}

	l, m := ParseLines(data, func(l Line) bool {
		return l.Start.X == l.End.X || l.Start.Y == l.End.Y
	})

	f := InitField(m)

	DrawLines(l, f)

	spew.Dump(f)
	assert.Equal(t, 5, f.CountOverlapOver(2))
}
