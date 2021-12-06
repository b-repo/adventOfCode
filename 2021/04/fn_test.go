package _4

import (
	"2021/common/reader"
	"github.com/stretchr/testify/assert"
	"testing"
)

var draw = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

func TestParseBoard(t *testing.T) {
	d, err := reader.ReadStringFileAndSplitByLine("testdata/04.log")
	if err != nil {
		panic(err.Error())
	}

	b := Board{
		Rows: [][]int{
			{22, 13, 17, 11, 0},
			{8, 2, 23, 4, 24},
			{21, 9, 14, 16, 7},
			{6, 10, 3, 18, 5},
			{1, 12, 20, 15, 19},
		},
		Columns: [][]int{
			{22, 8, 21, 6, 1},
			{13, 2, 9, 10, 12},
			{17, 23, 14, 3, 20},
			{11, 4, 16, 18, 15},
			{0, 24, 7, 5, 19},
		},
		Sum: 300,
	}

	r := parseBoard(d, 0, 5)

	assert.Equal(t, b, r)
}

func TestGetWinnerBoardID(t *testing.T) {
	d, err := reader.ReadStringFileAndSplitByLine("testdata/04.log")
	if err != nil {
		panic(err.Error())
	}

	b := ParseBoards(d)
	s1 := GetWinnerBoardID(b, draw)

	assert.Equal(t, 4512, s1)

}

func TestGetLooserBoardID(t *testing.T) {
	d, err := reader.ReadStringFileAndSplitByLine("testdata/04.log")
	if err != nil {
		panic(err.Error())
	}

	b := ParseBoards(d)
	s1 := GetLooserBoardID(b, draw)

	assert.Equal(t, 1924, s1)

}
