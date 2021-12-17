package _17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHigherVelocityReached(t *testing.T) {
	target, err := ParseTarget("testdata/17.log")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 45, GetHigherYReached(target))
}

func TestGetStrongestVector(t *testing.T) {
	target, err := ParseTarget("testdata/17.log")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 112, len(GetPossibleVectors(target)))
}
