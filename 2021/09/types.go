package _9

type HeightMap [][]int

type Coordinates struct {
	X, Y int
}

type Basin []Coordinates

func (h HeightMap) LowPoints() []Coordinates {
	c := make([]Coordinates, 0, len(h)*(len(h[0])/2))

	for x := range h {
		for y := range h[x] {
			if h[x][y] < lowestAdjacentLocations(h, x, y) {
				c = append(c, Coordinates{x, y})
			}
		}
	}

	return c
}

func (c *Coordinates) RiskLevel(h HeightMap) int {
	return h[c.X][c.Y] + 1
}

func (c Coordinates) FindBasin(h HeightMap) Basin {
	return extendsBasin(h, c, map[Coordinates]bool{})
}

func (h HeightMap) Basins() []Basin {
	lowPoints := h.LowPoints()

	b := make([]Basin, 0, len(lowPoints))

	for i := range lowPoints {
		b = append(b, lowPoints[i].FindBasin(h))
	}

	return b
}

func (b Basin) Size() int {
	return len(b)
}
