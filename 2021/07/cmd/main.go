package main

import (
	_7 "2021/07"
	"fmt"
	"strconv"
)

func main() {
	s := _7.InitCrabsSubmarines("data/07.log")

	p := s.ClosestAlignment()
	f1 := s.FuelConsumptionTo(int(p))
	f2 := s.CheapestAlignment()

	fmt.Println("Part 1 : " + strconv.Itoa(f1))
	fmt.Println("Part 2 : " + strconv.Itoa(f2))
}
