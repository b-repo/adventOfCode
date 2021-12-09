package main

import (
	_9 "2021/09"
	"2021/common/reader"
	"fmt"
	"strconv"
)

func main() {
	d, err := reader.Read2DIntSliceFile("data/09.log")
	if err != nil {
		panic("unable to read file.")
	}

	h := _9.HeightMap(d)
	l := h.LowPoints()

	fmt.Println("Part 1 : " + strconv.Itoa(_9.EvaluateRisks(h, l)))
	fmt.Println("Part 2 : " + strconv.Itoa(_9.MultiplyBasinsSizes(_9.GetTop3LargestBasins(h.Basins()))))
}
