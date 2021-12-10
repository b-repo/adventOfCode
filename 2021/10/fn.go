package _10

import "math/rand"

func FindErrorInLines(s [][]string) ([]string, [][]string) {
	r := make([]string, 0)
	r2 := make([][]string, 0)

	for i := range s {
		c, p := FindErrorInLine(s[i], 0, nil)
		if c != "" {
			r = append(r, c)
		} else {
			m := make([]string, 0)
			for p != nil {
				m = append(m, p.Opposite())
				p = p.Previous
			}
			r2 = append(r2, m)
		}
	}

	return r, r2
}

func FindErrorInLine(s []string, index int, previous *Element) (string, *Element) {
	if index >= len(s) {
		return "", previous
	}

	e := Element{
		Character: s[index],
		Previous:  previous,
	}

	if e.IsOpener() {
		return FindErrorInLine(s, index+1, &e)
	}

	if e.Character == previous.Opposite() {
		if previous.Previous == nil {
			return FindErrorInLine(s, index+1, nil)
		}

		e.Character = previous.Previous.Character
		e.Previous = previous.Previous.Previous
		return FindErrorInLine(s, index+1, &e)
	}

	return s[index], previous

}

func CalculateSyntaxCheckerScore(e []string) int {
	n := 0

	for i := range e {
		switch e[i] {
		case ")":
			n += 3
		case "]":
			n += 57
		case "}":
			n += 1197
		case ">":
			n += 25137
		default:
			panic("unknown character")
		}
	}

	return n
}

func CalculateAutoCompleteScore(m [][]string) int {
	r := CalculateAutoCompleteScores(m)

	return r[len(r)/2]
}

func CalculateAutoCompleteScores(m [][]string) []int {
	r := make([]int, 0, len(m))

	for i := range m {
		n := 0
		for j := range m[i] {
			n *= 5
			switch m[i][j] {
			case ")":
				n++
			case "]":
				n += 2
			case "}":
				n += 3
			case ">":
				n += 4
			default:
				panic("unknown character")
			}
		}
		r = append(r, n)
	}

	return quickSort(r)
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}
