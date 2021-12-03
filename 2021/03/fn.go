package _3

import (
	"strconv"
	"strings"
)

func ReadEnergyConsumption(data2d [][]string) EnergyConsumption {
	var g string
	var e string

	for i := 0; i < len(data2d[0]); i++ {
		c0, c1 := count(data2d, i)
		if c0 > c1 {
			g += "0"
			e += "1"
		} else {
			e += "0"
			g += "1"
		}
	}

	gi, err := strconv.ParseInt(g, 2, 64)
	if err != nil {
		panic("Unable to convert GammaRate to int")
	}
	ei, err := strconv.ParseInt(e, 2, 64)
	if err != nil {
		panic("Unable to convert EpsilonRate to int")
	}

	return EnergyConsumption{
		GammaRate:   gi,
		EpsilonRate: ei,
	}
}

func ReadLifeSupportRating(data2d [][]string) LifeSupportRating {
	o2, err := strconv.ParseInt(strings.Join(getO2GeneratorRate(data2d, 0), ""), 2, 64)
	if err != nil {
		panic("unable to convert O2GeneratorRate to int")
	}
	co2, err := strconv.ParseInt(strings.Join(getCO2ScrubberRate(data2d, 0), ""), 2, 64)
	if err != nil {
		panic("unable to convert CO2ScrubberRate to int")
	}

	return LifeSupportRating{
		O2GeneratorRate: o2,
		CO2ScrubberRate: co2,
	}
}

func count(data2d [][]string, index int) (float64, float64) {
	c0 := 0.
	c1 := 0.
	for j := 0; j < len(data2d); j++ {
		if data2d[j][index] == "1" {
			c1++
		} else {
			c0++
		}
	}

	return c0, c1
}

func getO2GeneratorRate(working2d [][]string, c int) []string {
	if len(working2d) == 2 {
		if working2d[0][len(working2d[0])-1] == "1" {
			return working2d[0]
		}
		return working2d[1]
	}

	if len(working2d) == 1 {
		return working2d[0]
	}

	m := "0"

	c0, c1 := count(working2d, c)
	if c0 < c1 {
		m = "1"
	}

	d := getDataThatStartWith(working2d, m, c)

	return getO2GeneratorRate(d, c+1)
}

func getCO2ScrubberRate(working2d [][]string, c int) []string {
	if len(working2d) == 2 {
		if working2d[0][len(working2d[0])-1] == "1" {
			return working2d[1]
		}
		return working2d[0]
	}

	if len(working2d) == 1 {
		return working2d[0]
	}

	m := "1"

	c0, c1 := count(working2d, c)
	if c0 < c1 {
		m = "0"
	}

	d := getDataThatStartWith(working2d, m, c)

	return getCO2ScrubberRate(d, c+1)
}

func getDataThatStartWith(d [][]string, major string, c int) [][]string {
	r := make([][]string, 0)

	for i := range d {
		if d[i][c] == major {
			r = append(r, d[i])
		}
	}

	return r
}
