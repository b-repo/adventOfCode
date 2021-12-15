package _15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseWholeCavernFromFile(t *testing.T) {
	g, m, l, err := ParseWholeCavernFromFile("testdata/15.log")
	if err != nil {
		t.Fail()
	}

	s := m[Coordinate{}]
	e := m[Coordinate{X: l - 1, Y: l - 1}]

	p, err := g.Shortest(s, e)

	assert.Equal(t, int64(315), p.Distance)
}

func TestShortest(t *testing.T) {
	g, m, l, err := ParseCavernFromFile("testdata/15.log")
	if err != nil {
		t.Fail()
	}

	s := m[Coordinate{}]
	e := m[Coordinate{X: l - 1, Y: l - 1}]

	p, err := g.Shortest(s, e)

	assert.Equal(t, int64(40), p.Distance)

	g, m, l, err = ParseCavernFromFile("testdata/15_big.log")
	if err != nil {
		t.Fail()
	}

	s = m[Coordinate{}]
	e = m[Coordinate{X: l - 1, Y: l - 1}]

	p, err = g.Shortest(s, e)

	assert.Equal(t, int64(315), p.Distance)
}
