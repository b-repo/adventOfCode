package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = []string{
	"199",
	"200",
	"208",
	"210",
	"200",
	"207",
	"240",
	"269",
	"260",
	"263",
}

func TestIncreasedAmount(t *testing.T) {
	assert.Equal(t, 7, increasedAmount(data))
}

func TestIncreasedAmountThreeMeasurement(t *testing.T) {
	assert.Equal(t, 5, increasedAmountThreeMeasurement(data))
}
