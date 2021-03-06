package _14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertPairs(t *testing.T) {
	tpl, m := ParsePolymerDataFromFile("testdata/14.log")
	assert.Equal(t, map[string]int{"NC": 1, "CN": 1, "NB": 1, "BC": 1, "CH": 1, "HB": 1}, InsertPairs(tpl, m, 1))
}

func TestFindLeastAndMostCommonElementInPolymer(t *testing.T) {
	tpl, m := ParsePolymerDataFromFile("testdata/14.log")

	mint := InsertPairs(tpl, m, 40)

	min, max := FindLeastAndMostCommonElementInPolymer(mint)

	assert.Equal(t, 3849876073, min)
	assert.Equal(t, 2192039569602, max)
}
