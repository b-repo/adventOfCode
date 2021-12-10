package main

import (
	_10 "2021/10"
	"2021/common/reader"
	"fmt"
	"strconv"
)

func main() {
	d, err := reader.Read2DStringSliceFile("data/10.log")
	if err != nil {
		panic("unable to read file.")
	}

	errors, missings := _10.FindErrorInLines(d)

	fmt.Println("Part 1 : " + strconv.Itoa(_10.CalculateSyntaxCheckerScore(errors)))
	fmt.Println("Part 2 : " + strconv.Itoa(_10.CalculateAutoCompleteScore(missings)))
}
