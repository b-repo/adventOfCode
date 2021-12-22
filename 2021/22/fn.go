package _22

import (
	"2021/common/reader"
	"strconv"
	"strings"
)

func ParseCore(fname string, f func(Cuboid) bool) (Core, error) {
	s, err := reader.ReadStringFileAndSplitByLine(fname)
	if err != nil {
		return nil, err
	}

	core := make([]Cuboid, 0, len(s))

	for _, line := range s {
		stateRange := strings.Split(line, " ")
		xMin, yMin, zMin, xMax, yMax, zMax, err2 := extractMinMaxPoints(stateRange[1])
		if err2 != nil {
			return nil, err2
		}

		c := NewCuboid(xMin, xMax, yMin, yMax, zMin, zMax, stateRange[0] == "on")
		if f == nil || f(c) {
			core = append(core, c)
		}
	}

	return core, nil
}

func extractMinMaxPoints(ptsRange string) (xMin, yMin, zMin, xMax, yMax, zMax int64, err error) {
	points := strings.Split(ptsRange, ",")
	xRange := strings.Split(strings.Replace(points[0], "x=", "", 1), "..")
	yRange := strings.Split(strings.Replace(points[1], "y=", "", 1), "..")
	zRange := strings.Split(strings.Replace(points[2], "z=", "", 1), "..")
	xMin, err = strconv.ParseInt(xRange[0], 10, 64)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, err
	}

	yMin, err = strconv.ParseInt(yRange[0], 10, 64)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, err
	}

	zMin, err = strconv.ParseInt(zRange[0], 10, 64)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, err
	}

	xMax, err = strconv.ParseInt(xRange[1], 10, 64)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, err
	}

	yMax, err = strconv.ParseInt(yRange[1], 10, 64)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, err
	}

	zMax, err = strconv.ParseInt(zRange[1], 10, 64)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, err
	}

	return
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}

	return a
}
