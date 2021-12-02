package _1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = []int{
	199,
	200,
	208,
	210,
	200,
	207,
	240,
	269,
	260,
	263,
}

func TestIncreasedAmount(t *testing.T) {
	assert.Equal(t, 7, IncreasedAmount(data))
}

func TestIncreasedAmountThreeMeasurement(t *testing.T) {
	assert.Equal(t, 5, IncreasedAmountThreeMeasurement(data))
}
