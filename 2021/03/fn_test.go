package _3

import (
	"2021/common/reader"
	"github.com/stretchr/testify/assert"
	"testing"
)

var nrj = EnergyConsumption{
	GammaRate:   22,
	EpsilonRate: 9,
}

var life = LifeSupportRating{
	O2GeneratorRate: 23,
	CO2ScrubberRate: 10,
}

func TestReadEnergyConsumption(t *testing.T) {
	d, err := reader.Read2DStringSliceFile("testdata/03.log")
	if err != nil {
		panic(err.Error())
	}

	assert.Equal(t, nrj, ReadEnergyConsumption(d))
}

func TestCount(t *testing.T) {
	d, err := reader.Read2DStringSliceFile("testdata/03.log")
	if err != nil {
		panic(err.Error())
	}

	assert.True(t, count(d, 0))
}

func TestReadLifeSupportRating(t *testing.T) {
	d, err := reader.Read2DStringSliceFile("testdata/03.log")
	if err != nil {
		panic(err.Error())
	}

	assert.Equal(t, life, ReadLifeSupportRating(d))
}

func TestGetDataThatStartWith(t *testing.T) {
	d := [][]string{
		{"0", "1", "0"},
		{"1", "1", "0"},
		{"0", "1", "1"},
	}

	assert.Equal(t, [][]string{{"1", "1", "0"}}, getDataThatStartWith(d, "1", 0))
	assert.Equal(t, [][]string{{"0", "1", "0"}, {"0", "1", "1"}}, getDataThatStartWith(d, "0", 0))
	assert.Equal(t, [][]string{}, getDataThatStartWith(d, "0", 1))
	assert.Equal(t, [][]string{{"0", "1", "1"}}, getDataThatStartWith(d, "1", 2))
}

func TestGetO2GeneratorRate(t *testing.T) {
	d, err := reader.Read2DStringSliceFile("testdata/03.log")
	if err != nil {
		panic(err.Error())
	}

	assert.Equal(t, []string{"1", "0", "1", "1", "1"}, getO2GeneratorRate(d, 0))
}

func TestGetCO2ScrubberRate(t *testing.T) {
	d, err := reader.Read2DStringSliceFile("testdata/03.log")
	if err != nil {
		panic(err.Error())
	}

	assert.Equal(t, []string{"0", "1", "0", "1", "0"}, getCO2ScrubberRate(d, 0))
}
