package _7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrabsSubmarines_ClosestAlignment(t *testing.T) {
	s := InitCrabsSubmarines("testdata/07.log")

	assert.Equal(t, 2., s.ClosestAlignment())
}

func TestCrabsSubmarines_FuelConsumptionTo(t *testing.T) {
	s := InitCrabsSubmarines("testdata/07.log")

	assert.Equal(t, 37, s.FuelConsumptionTo(2))
}

func TestCrabsSubmarines_CheapestAlignment(t *testing.T) {
	s := InitCrabsSubmarines("testdata/07.log")

	assert.Equal(t, 168, s.CheapestAlignment())
}
