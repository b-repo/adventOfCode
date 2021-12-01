package main

import (
	"2021/2021/common/reader"
	"fmt"
)

func main() {
	data, err := reader.ReadIntFileAndSplitByLine("2021/data/01.log")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(increasedAmount(data))
	fmt.Println(increasedAmountThreeMeasurement(data))
}

func increasedAmount(d []int) int {
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

func increasedAmountThreeMeasurement(d []int) int {
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
