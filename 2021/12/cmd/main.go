package main

import (
	_12 "2021/12"
	"fmt"
	"strconv"
)

func main() {
	sc, err := _12.ReadCavesMapFromFile("data/12.log")
	if err != nil {
		panic("unable to read file.")
	}

	paths1 := _12.Explore(sc, "start", "start")
	paths2 := _12.ExploreTwice(sc, "start", "start")

	fmt.Println("Part 1 : " + strconv.Itoa(len(paths1)))
	fmt.Println("Part 2 : " + strconv.Itoa(len(paths2)))
}
