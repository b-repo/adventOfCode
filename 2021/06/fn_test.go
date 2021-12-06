package _6

import (
	"2021/common/reader"
	"github.com/stretchr/testify/assert"
	"testing"
)

var expectedSchool = SchoolOfFish{0, 1, 1, 2, 1, 0, 0, 0, 0}

func TestInitSchoolOfFish(t *testing.T) {
	data, err := reader.ReadIntFileAndSplitBy("testdata/06.log", ",")
	if err != nil {
		panic(err.Error())
	}

	s := InitSchoolOfFish(data)

	assert.Equal(t, expectedSchool, s)
}

func TestSchoolOfFish_LiveAnotherDay(t *testing.T) {
	data, err := reader.ReadIntFileAndSplitBy("testdata/06.log", ",")
	if err != nil {
		panic(err.Error())
	}

	s := InitSchoolOfFish(data)

	for i := 0; i < 80; i++ {
		s.LiveAnotherDay()
	}

	assert.Equal(t, 5934, s.Population())
}
