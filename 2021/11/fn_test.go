package _11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterate(t *testing.T) {
	b, err := ReadOctopusGridFromFile("testdata/11.log")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 1656, Iterate(100, b))
}

func TestIterateUntilBigFlash(t *testing.T) {
	b, err := ReadOctopusGridFromFile("testdata/11.log")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 195, IterateUntilBigFlash(b))
}
