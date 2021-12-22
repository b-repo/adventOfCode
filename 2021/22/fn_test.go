package _22

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCore(t *testing.T) {
	c, err := ParseCore("testdata/22_small.log", func(c Cuboid) bool { return !c.isOn })
	if err != nil {
		t.Fail()
	}

	expectedCore := Core{{
		isOn: false,
		Min:  Coordinate{9, 9, 9},
		Max:  Coordinate{11, 11, 11},
	}}

	assert.Equal(t, expectedCore, c)
}

func TestProcessInstruction(t *testing.T) {
	c, err := ParseCore("testdata/22_small.log", func(cuboid Cuboid) bool {
		return cuboid.Min.X >= -50 && cuboid.Min.Y >= -50 && cuboid.Min.Z >= -50 &&
			cuboid.Max.X <= 50 && cuboid.Max.Y <= 50 && cuboid.Max.Z <= 50
	})
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, int64(39), c.Compute())

	c, err = ParseCore("testdata/22.log", func(cuboid Cuboid) bool {
		return cuboid.Min.X >= -50 && cuboid.Min.Y >= -50 && cuboid.Min.Z >= -50 &&
			cuboid.Max.X <= 50 && cuboid.Max.Y <= 50 && cuboid.Max.Z <= 50
	})
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, int64(590784), c.Compute())
}
