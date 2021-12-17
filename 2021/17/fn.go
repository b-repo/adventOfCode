package _17

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ParseTarget(fname string) (Target, error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return Target{}, err
	}

	s := strings.Split(strings.Replace(string(b), "target area: ", "", 1), ", ")

	xRange := strings.Split(strings.Replace(s[0], "x=", "", 1), "..")
	yRange := strings.Split(strings.Replace(s[1], "y=", "", 1), "..")

	xMin, err := strconv.Atoi(xRange[0])
	if err != nil {
		return Target{}, err
	}

	xMax, err := strconv.Atoi(xRange[1])
	if err != nil {
		return Target{}, err
	}

	yMin, err := strconv.Atoi(yRange[0])
	if err != nil {
		return Target{}, err
	}

	yMax, err := strconv.Atoi(yRange[1])
	if err != nil {
		return Target{}, err
	}

	return Target{
		Xmin: xMin,
		Xmax: xMax,
		Ymin: yMin,
		Ymax: yMax,
	}, nil
}

func GetHigherYReached(t Target) int {
	min := -t.Ymin - 1

	return min * (min + 1) / 2
}

func GetPossibleVectors(t Target) []Vector {
	vectors := make([]Vector, 0)

	for x := t.Xmin; x <= t.Xmax; x++ {
		for y := t.Ymin; y <= t.Ymax; y++ {
			vectors = append(vectors, Vector{
				X: x,
				Y: y,
			})
		}
	}

	for y := t.Ymax; y <= -1*t.Ymin-1; y++ {
		for x := 0; x < t.Xmin; x++ {
			v := Vector{x, y}

			if v.WillReachTarget(t) {
				vectors = append(vectors, v)
			}
		}
	}

	return vectors
}
