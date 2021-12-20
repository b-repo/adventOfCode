package _20

import (
	"math"
)

type Algorithm string

type Coordinates struct {
	X, Y int
}

type Pixels map[Coordinates]bool

func (p Pixels) GetBorders() (int, int, int, int) {
	var maxX, maxY int
	var minX, minY = math.MaxInt, math.MaxInt

	for c := range p {
		if c.X > maxX {
			maxX = c.X
		}

		if c.Y > maxY {
			maxY = c.Y
		}

		if c.X < minX {
			minX = c.X
		}

		if c.Y < minY {
			minY = c.Y
		}
	}

	return minX, minY, maxX, maxY
}

type Image struct {
	Image           Pixels
	LitPixels, Step int
}

func (img *Image) Enhance(a Algorithm) {
	enhancement := make(Pixels)
	newLitPixels := 0

	minX, minY, maxX, maxY := img.Image.GetBorders()

	for i := minX - 1; i <= maxX+1; i++ {
		for j := minY - 1; j <= maxY+1; j++ {
			aIndex := 0

			for _, x := range []int{i - 1, i, i + 1} {
				for _, y := range []int{j - 1, j, j + 1} {
					aIndex = aIndex << 1
					if pixel, ok := img.Image[Coordinates{x, y}]; ok {
						if pixel {
							aIndex |= 1
						}
					} else if a[0] == '#' && img.Step%2 == 1 {
						aIndex |= 1
					}
				}
			}

			enhancement[Coordinates{i, j}] = a[aIndex] == '#'
			if a[aIndex] == '#' {
				newLitPixels++
			}
		}
	}

	img.Step++
	img.Image = enhancement
	img.LitPixels = newLitPixels
}

func (img *Image) String() string {
	s := ""
	minX, minY, maxX, maxY := img.Image.GetBorders()
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if img.Image[Coordinates{x, y}] {
				s += "#"

				continue
			}
			s += "."
		}
		s += "\n"
	}

	return s
}
