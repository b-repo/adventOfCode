package _7

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type CrabsSubmarines struct {
	Position []int
	Max      int
	Min      int
	Moy      float64
}

func InitCrabsSubmarines(fname string) CrabsSubmarines {
	max := 0
	min := math.MaxInt
	sum := 0

	b, err := ioutil.ReadFile(fname)
	if err != nil {
		panic("unable to parse file")
	}

	position := strings.Split(string(b), ",")
	// Assign cap to avoid resize on every append.
	nums := make([]int, 0, len(position))

	for i := range position {
		n, err := strconv.Atoi(position[i])
		if err != nil {
			panic("unable to convert in integer")
		}

		sum += n

		if max < n {
			max = n
		}

		if min > n {
			min = n
		}

		nums = append(nums, n)
	}

	return CrabsSubmarines{
		Position: nums,
		Max:      max,
		Min:      min,
		Moy:      float64(sum) / float64(len(nums)),
	}
}

func (c CrabsSubmarines) CheapestAlignment() int {
	m := math.MaxInt
	for i := c.Min; i <= c.Max; i++ {
		s := 0
		for j := range c.Position {
			s += (Abs(c.Position[j]-i) * (Abs(c.Position[j]-i) + 1)) / 2
		}
		if s < m {
			m = s
		}
	}

	return m
}

func (c CrabsSubmarines) ClosestAlignment() float64 {
	med := Dichotome(c.Position, 0, c.Moy)

	return math.Round(med)
}

func (c CrabsSubmarines) FuelConsumptionTo(position int) int {
	f := 0

	for i := range c.Position {
		m := c.Position[i] - position
		if m < 0 {
			m *= -1
		}
		f += m
	}

	return f
}
