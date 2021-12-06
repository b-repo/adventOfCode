package _5

import (
	"math"
	"strconv"
	"strings"
)

type Field struct {
	Map        [][]int
	Overlap    map[int]int
	MaxOverlap int
}

type Line struct {
	Start Coordinate
	End   Coordinate
}

type Coordinate struct {
	X int
	Y int
}

func (f *Field) CountOverlapOver(n int) int {
	s := 0
	for i := n; i <= f.MaxOverlap; i++ {
		s += f.Overlap[i]
	}
	return s
}

func (l *Line) Draw(f *Field) {
	var absA, absO, length float64

	a := float64(l.End.X - l.Start.X)
	o := float64(l.End.Y - l.Start.Y)

	if a != 0 {
		absA = math.Abs(a)
		length = absA
	}

	if o != 0 {
		absO = math.Abs(o)
		length = absO
	}

	if absA != 0 && absO != 0 && absA != absO {
		panic("unable to draw this diagonal")
	}

	for c := 0.; c <= length; c++ {
		var offsetA, offsetO int

		if absA != 0 {
			offsetA = int((a / absA) * c)
		}

		if absO != 0 {
			offsetO = int((o / absO) * c)
		}

		f.Map[l.Start.X+offsetA][l.Start.Y+offsetO] += 1

		_, ok := f.Overlap[f.Map[l.Start.X+offsetA][l.Start.Y+offsetO]]
		if !ok {
			f.Overlap[f.Map[l.Start.X+offsetA][l.Start.Y+offsetO]] = 1
			f.MaxOverlap = f.Map[l.Start.X+offsetA][l.Start.Y+offsetO]
		} else {
			f.Overlap[f.Map[l.Start.X+offsetA][l.Start.Y+offsetO]]++
		}

		if f.Map[l.Start.X+offsetA][l.Start.Y+offsetO] > 1 {
			f.Overlap[(f.Map[l.Start.X+offsetA][l.Start.Y+offsetO])-1]--
		}
	}
}

func ParseLine(s string) (Line, int) {
	max := 0

	ends := strings.Split(s, " -> ")

	startCoordinateStr := strings.Split(ends[0], ",")

	endCoordinateStr := strings.Split(ends[1], ",")

	startX, err := strconv.Atoi(startCoordinateStr[0])
	if err != nil {
		panic("unable to parse int of startX.")
	}

	if startX > max {
		max = startX
	}

	startY, err := strconv.Atoi(startCoordinateStr[1])
	if err != nil {
		panic("unable to parse int of startY.")
	}

	if startY > max {
		max = startY
	}

	endX, err := strconv.Atoi(endCoordinateStr[0])
	if err != nil {
		panic("unable to parse int of endX.")
	}

	if endX > max {
		max = endX
	}

	endY, err := strconv.Atoi(endCoordinateStr[1])
	if err != nil {
		panic("unable to parse int of endY.")
	}

	if endY > max {
		max = endY
	}

	return Line{
		Start: Coordinate{
			X: startX,
			Y: startY,
		},
		End: Coordinate{
			X: endX,
			Y: endY,
		},
	}, max
}

func InitField(size int) *Field {
	m := make([][]int, 0, size+1)

	for i := 0; i <= size; i++ {
		c := make([]int, 0, size+1)
		for j := 0; j <= size; j++ {
			c = append(c, 0)
		}
		m = append(m, c)
	}

	return &Field{
		Map:        m,
		Overlap:    map[int]int{},
		MaxOverlap: 0,
	}
}
