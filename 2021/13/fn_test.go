package _13

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTransparentPaperFromFile(t *testing.T) {
	d, err := ParseTransparentPaperFromFile("testdata/13.log")
	if err != nil {
		t.Fail()
	}

	tp := &TransparentPaper{
		Grid: [][]bool{
			{false, false, false, true, false, false, true, false, false, true, false},
			{false, false, false, false, true, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false, false, false},
			{true, false, false, false, false, false, false, false, false, false, false},
			{false, false, false, true, false, false, false, false, true, false, true},
			{false, false, false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false, false, false},
			{false, false, false, false, false, false, false, false, false, false, false},
			{false, true, false, false, false, false, true, false, true, true, false},
			{false, false, false, false, true, false, false, false, false, false, false},
			{false, false, false, false, false, false, true, false, false, false, true},
			{true, false, false, false, false, false, false, false, false, false, false},
			{true, false, true, false, false, false, false, false, false, false, false},
		},
		Instructions: []Instruction{
			{IsHorizontalFold: true, Value: 7},
			{IsHorizontalFold: false, Value: 5},
		},
	}

	assert.Equal(t, tp, d)

}

func TestTransparentPaper_FoldHorizontally(t *testing.T) {
	d, err := ParseTransparentPaperFromFile("testdata/13.log")
	if err != nil {
		t.Fail()
	}

	expectedGrid := [][]bool{
		{true, false, true, true, false, false, true, false, false, true, false},
		{true, false, false, false, true, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false, false, false, true},
		{true, false, false, false, true, false, false, false, false, false, false},
		{false, true, false, true, false, false, true, false, true, true, true},
		{false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false},
	}

	d.FoldHorizontally(7)

	assert.Equal(t, expectedGrid, d.Grid)

	expectedGrid = [][]bool{
		{true, false, true, true, false, false, true, false, false, true, false},
		{true, false, false, false, true, false, false, false, false, false, false},
		{false, true, false, true, false, false, true, false, true, true, true},
	}

	d.FoldHorizontally(3)

	assert.Equal(t, expectedGrid, d.Grid)
}

func TestTransparentPaper_FoldVertically(t *testing.T) {
	d := TransparentPaper{Grid: [][]bool{
		{true, false, true, true, false, false, true, false, false, true, false},
		{true, false, false, false, true, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false, false, false, true},
		{true, false, false, false, true, false, false, false, false, false, false},
		{false, true, false, true, false, false, true, false, true, true, true},
		{false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false},
	}}

	expectedGrid := [][]bool{
		{true, true, true, true, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, false, false, false, true},
		{true, true, true, true, true},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}

	d.FoldVertically(5)

	assert.Equal(t, expectedGrid, d.Grid)
}
