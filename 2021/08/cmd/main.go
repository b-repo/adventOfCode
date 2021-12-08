package main

import (
	_8 "2021/08"
	"fmt"
	"strconv"
)

func main() {
	s, err := _8.ReadSignalsAndDigits("data/08.log")
	if err != nil {
		panic("unable to read file.")
	}

	fmt.Println("Part 1 : " + strconv.Itoa(_8.CountDigitsWithUniqueSegmentsAmount(s)))
	fmt.Println("Part 2 : " + strconv.Itoa(_8.SumDrecryptedDigits(s)))
}
