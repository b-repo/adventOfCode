package _9

import "math"

func lowestAdjacentLocations(h HeightMap, x, y int) int {
	l := math.MaxInt

	if x > 0 && h[x-1][y] < l {
		l = h[x-1][y]
	}

	if x < len(h)-1 && h[x+1][y] < l {
		l = h[x+1][y]
	}

	if y > 0 && h[x][y-1] < l {
		l = h[x][y-1]
	}

	if y < len(h[x])-1 && h[x][y+1] < l {
		l = h[x][y+1]
	}

	return l
}

func extendsBasin(h HeightMap, c Coordinates, v map[Coordinates]bool) []Coordinates {
	coordinates := make([]Coordinates, 0)

	visited, found := v[c]
	if found && visited {
		return coordinates
	}

	v[c] = true
	coordinates = append(coordinates, c)

	n := h[c.X][c.Y]

	if c.X > 0 && h[c.X-1][c.Y] != 9 && h[c.X-1][c.Y] > n {
		coordinates = append(coordinates, extendsBasin(h, Coordinates{c.X - 1, c.Y}, v)...)
	}

	if c.X < len(h)-1 && h[c.X+1][c.Y] != 9 && h[c.X+1][c.Y] > n {
		coordinates = append(coordinates, extendsBasin(h, Coordinates{c.X + 1, c.Y}, v)...)
	}

	if c.Y > 0 && h[c.X][c.Y-1] != 9 && h[c.X][c.Y-1] > n {
		coordinates = append(coordinates, extendsBasin(h, Coordinates{c.X, c.Y - 1}, v)...)
	}

	if c.Y < len(h[c.X])-1 && h[c.X][c.Y+1] != 9 && h[c.X][c.Y+1] > n {
		coordinates = append(coordinates, extendsBasin(h, Coordinates{c.X, c.Y + 1}, v)...)
	}

	return coordinates
}

func EvaluateRisks(h HeightMap, lowestPoints []Coordinates) int {
	a := 0

	for l := range lowestPoints {
		a += lowestPoints[l].RiskLevel(h)
	}

	return a
}

func GetTop3LargestBasins(b []Basin) []Basin {
	top1 := Basin{}
	top2 := Basin{}
	top3 := Basin{}

	for i := range b {
		if top1.Size() == 0 || b[i].Size() > top1.Size() {
			top3 = top2
			top2 = top1
			top1 = b[i]
			continue
		}
		if top2.Size() == 0 || b[i].Size() > top2.Size() {
			top3 = top2
			top2 = b[i]
			continue
		}
		if top3.Size() == 0 || b[i].Size() > top3.Size() {
			top3 = b[i]
			continue
		}
	}

	return []Basin{top1, top2, top3}
}

func MultiplyBasinsSizes(b []Basin) int {
	n := 1

	for i := range b {
		n *= b[i].Size()
	}

	return n
}
