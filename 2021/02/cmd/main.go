package main

import (
	_2 "2021/02"
	"2021/common/reader"
	"fmt"
	"strconv"
)

func main() {
	data, err := reader.ReadStringFileAndSplitByLine("data/02.log")
	if err != nil {
		panic(err.Error())
	}

	commands := _2.ReadCommandsFromData(data)

	h1, d1 := _2.CalculateFinalePosition(commands)
	h2, d2 := _2.CalculateFinalePositionWithAim(commands)

	fmt.Println("Part 1 : " + strconv.Itoa(h1*d1))
	fmt.Println("Part 2 : " + strconv.Itoa(h2*d2))
}
