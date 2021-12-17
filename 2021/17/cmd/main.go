package main

import (
	_17 "2021/17"
	"fmt"
	"log"
)

func main() {
	t, err := _17.ParseTarget("data/17.log")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Part 1 : %d\n", _17.GetHigherYReached(t))
	fmt.Printf("Part 2 : %d\n", len(_17.GetPossibleVectors(t)))

}
