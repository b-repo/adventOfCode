package _15

import (
	"2021/common/reader"
	"github.com/RyanCarrier/dijkstra"
)

func ParseCavernFromFile(fname string) (*dijkstra.Graph, map[Coordinate]int, int, error) {
	d, err := reader.Read2DIntSliceFile(fname)
	if err != nil {
		return nil, nil, 0, err
	}

	g := dijkstra.NewGraph()
	m := map[Coordinate]int{}
	n := 0

	for i := range d {
		for j := range d[i] {
			m[Coordinate{i, j}] = n
			g.AddVertex(n)
			n++
		}
	}

	for i := range d {
		for j := range d[i] {
			if i > 0 {
				g.AddArc(m[Coordinate{i - 1, j}], m[Coordinate{i, j}], int64(d[i][j]))
			}

			if i < len(d[i])-1 {
				g.AddArc(m[Coordinate{i + 1, j}], m[Coordinate{i, j}], int64(d[i][j]))
			}

			if j > 0 {
				g.AddArc(m[Coordinate{i, j - 1}], m[Coordinate{i, j}], int64(d[i][j]))
			}

			if j < len(d[i])-1 {
				g.AddArc(m[Coordinate{i, j + 1}], m[Coordinate{i, j}], int64(d[i][j]))
			}
		}
	}

	return g, m, len(d), nil
}

func ParseWholeCavernFromFile(fname string) (*dijkstra.Graph, map[Coordinate]int, int, error) {
	d, err := reader.Read2DIntSliceFile(fname)
	if err != nil {
		return nil, nil, 0, err
	}

	g := dijkstra.NewGraph()
	m := map[Coordinate]int{}
	n := 0

	for i := 0; i < 5*len(d); i++ {
		for j := 0; j < 5*len(d); j++ {
			m[Coordinate{i, j}] = n
			g.AddVertex(n)
			n++
		}
	}

	bigD := make([][]int, 0, len(d)*5)
	for i := 0; i < len(d)*5; i++ {
		bigD = append(bigD, make([]int, 0, len(d)*5))
	}

	for i := 0; i < len(d)*5; i++ {
		for j := 0; j < len(d)*5; j++ {
			if i < len(d) && j < len(d) {
				bigD[i] = append(bigD[i], d[i][j])

				continue
			}

			if i >= len(d) && j < len(d) {
				v := bigD[i%len(d)][j]
				v += i / len(d)
				if v > 9 {
					v = v % 9
				}
				bigD[i] = append(bigD[i], v)

				continue
			}

			if j >= len(d) && i < len(d) {
				v := bigD[i][j%len(d)]
				v += j / len(d)
				if v > 9 {
					v = v % 9
				}
				bigD[i] = append(bigD[i], v)

				continue
			}

			v := bigD[i%len(d)][j%len(d)]
			v += i/len(d) + j/len(d)
			if v > 9 {
				v = v % 9
			}
			bigD[i] = append(bigD[i], v)
		}
	}

	for i := range bigD {
		for j := range bigD[i] {
			if i > 0 {
				g.AddArc(m[Coordinate{i - 1, j}], m[Coordinate{i, j}], int64(bigD[i][j]))
			}

			if i < len(bigD[i])-1 {
				g.AddArc(m[Coordinate{i + 1, j}], m[Coordinate{i, j}], int64(bigD[i][j]))
			}

			if j > 0 {
				g.AddArc(m[Coordinate{i, j - 1}], m[Coordinate{i, j}], int64(bigD[i][j]))
			}

			if j < len(bigD[i])-1 {
				g.AddArc(m[Coordinate{i, j + 1}], m[Coordinate{i, j}], int64(bigD[i][j]))
			}
		}
	}

	return g, m, len(bigD), nil
}
