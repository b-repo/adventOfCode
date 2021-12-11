package main

import (
	_11 "2021/11"
	"fmt"
	"strconv"
)

func main() {
	d, err := _11.ReadOctopusGridFromFile("data/11.log")
	if err != nil {
		panic("unable to read file.")
	}

	dbf, err := _11.ReadOctopusGridFromFile("data/11.log")
	if err != nil {
		panic("unable to read file.")
	}

	flashes := _11.Iterate(100, d)
	firstBigFlash := _11.IterateUntilBigFlash(dbf)

	fmt.Println("Part 1 : " + strconv.Itoa(flashes))
	fmt.Println("Part 2 : " + strconv.Itoa(firstBigFlash))
}
