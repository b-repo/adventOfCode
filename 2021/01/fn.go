// Package _1 contains the fn used to resolve puzzle 01 from adventOfCode2021
package _1

func IncreasedAmount(d []int) int {
	a := 0

	var n, p int

	for i := 1; i < len(d); i++ {
		n = d[i]
		p = d[i-1]
		if n > p {
			a++
		}
	}

	return a
}

func IncreasedAmountThreeMeasurement(d []int) int {
	a := 0

	var n1, n2, n3, p1, p2, p3 int
	for i := 0; i < len(d)-3; i++ {
		n1 = d[i]
		n2 = d[i+1]
		n3 = d[i+2]
		p1 = d[i+1]
		p2 = d[i+2]
		p3 = d[i+3]
		if n1+n2+n3 < p1+p2+p3 {
			a++
		}
	}

	return a
}
