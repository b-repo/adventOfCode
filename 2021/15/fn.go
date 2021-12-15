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

	addVertex(len(d), m, g)

	graph, m2, i, err2 := addArcs(d, g, m)
	if err2 != nil {
		return graph, m2, i, err2
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

	addVertex(5*len(d), m, g)

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

	graph, m2, i, err2 := addArcs(bigD, g, m)
	if err2 != nil {
		return graph, m2, i, err2
	}

	return g, m, len(bigD), nil
}

func addVertex(len int, m map[Coordinate]int, g *dijkstra.Graph) {
	n := 0
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			m[Coordinate{i, j}] = n
			g.AddVertex(n)
			n++
		}
	}
}

func addArcs(d [][]int, g *dijkstra.Graph, m map[Coordinate]int) (*dijkstra.Graph, map[Coordinate]int, int, error) {
	for i := range d {
		for j := range d[i] {
			if i > 0 {
				err := g.AddArc(m[Coordinate{i - 1, j}], m[Coordinate{i, j}], int64(d[i][j]))
				if err != nil {
					return nil, nil, 0, err
				}
			}

			if i < len(d[i])-1 {
				err := g.AddArc(m[Coordinate{i + 1, j}], m[Coordinate{i, j}], int64(d[i][j]))
				if err != nil {
					return nil, nil, 0, err
				}
			}

			if j > 0 {
				err := g.AddArc(m[Coordinate{i, j - 1}], m[Coordinate{i, j}], int64(d[i][j]))
				if err != nil {
					return nil, nil, 0, err
				}
			}

			if j < len(d[i])-1 {
				err := g.AddArc(m[Coordinate{i, j + 1}], m[Coordinate{i, j}], int64(d[i][j]))
				if err != nil {
					return nil, nil, 0, err
				}
			}
		}
	}
	return nil, nil, 0, nil
}
