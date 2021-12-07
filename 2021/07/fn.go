package _7

import "math"

func Abs(i int) int {
	if i < 0 {
		return i * (-1)
	}
	return i
}

func Dichotome(s []int, previous, current float64) float64 {
	if int(previous) == int(current) {
		return current
	}

	lt := count(s, func(i int) bool { return float64(i) < current })
	gt := len(s) - lt

	if lt > gt {
		return Dichotome(s, current, current-(math.Abs(previous-current)/2))
	}

	if gt > lt {
		return Dichotome(s, current, current+(math.Abs(previous-current)/2))
	}

	return current
}

func count(s []int, f func(int) bool) int {
	a := 0

	for i := range s {
		if f(s[i]) {
			a++
		}
	}

	return a
}
