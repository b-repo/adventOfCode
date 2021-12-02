package _2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var data = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

var commands = []Command{
	{
		Direction: "forward",
		Power:     5,
	},
	{
		Direction: "down",
		Power:     5,
	},
	{
		Direction: "forward",
		Power:     8,
	},
	{
		Direction: "up",
		Power:     3,
	},
	{
		Direction: "down",
		Power:     8,
	},
	{
		Direction: "forward",
		Power:     2,
	},
}

func TestReadCommandsFromData(t *testing.T) {
	assert.Equal(t, commands, ReadCommandsFromData(data))
}

func TestCalculateFinalePosition(t *testing.T) {
	h, d := CalculateFinalePosition(commands)
	assert.Equal(t, h, 15)
	assert.Equal(t, d, 10)
}

func TestCalculateFinalePositionWithAim(t *testing.T) {
	h, d := CalculateFinalePositionWithAim(commands)
	assert.Equal(t, h, 15)
	assert.Equal(t, d, 60)
}
