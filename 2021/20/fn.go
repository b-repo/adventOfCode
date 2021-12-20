package _20

import (
	"io/ioutil"
	"strings"
)

func ParseInput(fname string) (Algorithm, *Image, error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return "", nil, err
	}

	d := strings.Split(string(b), "\n\n")
	rows := strings.Split(d[1], "\n")
	img := Image{Image: make(map[Coordinates]bool), LitPixels: 0, Step: 0}

	for x := range rows {
		for y, c := range rows[x] {
			img.Image[Coordinates{x, y}] = c == '#'
			if c == '#' {
				img.LitPixels++
			}
		}
	}

	return Algorithm(d[0]), &img, nil
}

func EnhanceImage(n int, img *Image, a Algorithm) {
	for i := 0; i < n; i++ {
		img.Enhance(a)
	}
}
