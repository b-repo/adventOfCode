package main

import (
	_14 "2021/14"
	"fmt"
	"strconv"
)

func main() {
	tpl10, m := _14.ParsePolymerDataFromFile("data/14.log")

	p10 := _14.InsertPairs(tpl10, m, 10)
	min10, max10 := _14.FindLeastAndMostCommonElementInPolymer(p10)

	tpl40, m := _14.ParsePolymerDataFromFile("data/14.log")

	p40 := _14.InsertPairs(tpl40, m, 40)
	min40, max40 := _14.FindLeastAndMostCommonElementInPolymer(p40)

	fmt.Println("Part 1 : " + strconv.Itoa(max10-min10))
	fmt.Println("Part 2 : " + strconv.Itoa(max40-min40))
}
