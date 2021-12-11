package _11

import (
	"2021/common/reader"
	"strconv"
	"strings"
)

func ReadOctopusGridFromFile(fname string) ([][]Octopus, error) {
	d, err := reader.ReadStringFileAndSplitByLine(fname)
	if err != nil {
		return nil, err
	}

	r := make([][]Octopus, 0)

	for i := range d {
		ys := strings.Split(d[i], "")
		col := make([]Octopus, 0)
		for j := range ys {
			n, err := strconv.Atoi(ys[j])
			if err != nil {
				panic(err.Error())
			}

			col = append(col, Octopus{Energy: n, X: i, Y: j})
		}

		r = append(r, col)
	}

	return r, nil
}

func Iterate(n int, b [][]Octopus) int {
	flashes := 0

	for i := 0; i < n; i++ {
		Next(b, &flashes)
	}

	return flashes
}

func IterateUntilBigFlash(b [][]Octopus) int {
	i := 0
	p := -1
	f := 0

	for {
		i++
		Next(b, &f)
		if f == p+len(b)*len(b[0]) {
			return i
		}
		p = f
	}
}

func Next(b [][]Octopus, flashes *int) {
	for i := range b {
		for j := range b[i] {
			b[i][j].HasFlashed = false
		}
	}

	for i := range b {
		for j := range b[i] {
			b[i][j].Gain(b, flashes)
		}
	}
}
