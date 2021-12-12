package _12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExplore(t *testing.T) {
	m1, err := ReadCavesMapFromFile("testdata/12_small.log")
	if err != nil {
		t.Fail()
	}

	expectedPaths_small := []string{
		"start,A,b,A,c,A,end",
		"start,A,b,A,end",
		"start,A,b,end",
		"start,A,c,A,b,A,end",
		"start,A,c,A,b,end",
		"start,A,c,A,end",
		"start,A,end",
		"start,b,A,c,A,end",
		"start,b,A,end",
		"start,b,end",
	}

	assert.ElementsMatch(t, expectedPaths_small, Explore(m1, "start", "start"))

	m2, err := ReadCavesMapFromFile("testdata/12_medium.log")
	if err != nil {
		t.Fail()
	}

	expectedPaths_medium := []string{
		"start,HN,dc,HN,end",
		"start,HN,dc,HN,kj,HN,end",
		"start,HN,dc,end",
		"start,HN,dc,kj,HN,end",
		"start,HN,end",
		"start,HN,kj,HN,dc,HN,end",
		"start,HN,kj,HN,dc,end",
		"start,HN,kj,HN,end",
		"start,HN,kj,dc,HN,end",
		"start,HN,kj,dc,end",
		"start,dc,HN,end",
		"start,dc,HN,kj,HN,end",
		"start,dc,end",
		"start,dc,kj,HN,end",
		"start,kj,HN,dc,HN,end",
		"start,kj,HN,dc,end",
		"start,kj,HN,end",
		"start,kj,dc,HN,end",
		"start,kj,dc,end",
	}

	assert.ElementsMatch(t, expectedPaths_medium, Explore(m2, "start", "start"))
}

func TestExploreTwice(t *testing.T) {
	m1, err := ReadCavesMapFromFile("testdata/12_small.log")
	if err != nil {
		t.Fail()
	}

	expectedSmall := []string{
		"start,A,b,A,b,A,c,A,end",
		"start,A,b,A,b,A,end",
		"start,A,b,A,b,end",
		"start,A,b,A,c,A,b,A,end",
		"start,A,b,A,c,A,b,end",
		"start,A,b,A,c,A,c,A,end",
		"start,A,b,A,c,A,end",
		"start,A,b,A,end",
		"start,A,b,d,b,A,c,A,end",
		"start,A,b,d,b,A,end",
		"start,A,b,d,b,end",
		"start,A,b,end",
		"start,A,c,A,b,A,b,A,end",
		"start,A,c,A,b,A,b,end",
		"start,A,c,A,b,A,c,A,end",
		"start,A,c,A,b,A,end",
		"start,A,c,A,b,d,b,A,end",
		"start,A,c,A,b,d,b,end",
		"start,A,c,A,b,end",
		"start,A,c,A,c,A,b,A,end",
		"start,A,c,A,c,A,b,end",
		"start,A,c,A,c,A,end",
		"start,A,c,A,end",
		"start,A,end",
		"start,b,A,b,A,c,A,end",
		"start,b,A,b,A,end",
		"start,b,A,b,end",
		"start,b,A,c,A,b,A,end",
		"start,b,A,c,A,b,end",
		"start,b,A,c,A,c,A,end",
		"start,b,A,c,A,end",
		"start,b,A,end",
		"start,b,d,b,A,c,A,end",
		"start,b,d,b,A,end",
		"start,b,d,b,end",
		"start,b,end",
	}

	assert.ElementsMatch(t, expectedSmall, ExploreTwice(m1, "start", "start"))
}
