package main

import (
	_15 "2021/15"
	"fmt"
	"log"
)

func main() {
	g, m, l, err := _15.ParseCavernFromFile("data/15.log")
	if err != nil {
		log.Fatal(err.Error())
	}

	s := m[_15.Coordinate{}]
	e := m[_15.Coordinate{X: l - 1, Y: l - 1}]

	p, err := g.Shortest(s, e)

	fmt.Printf("Part 1 : %d\n", p.Distance)

	g2, m2, l2, err := _15.ParseWholeCavernFromFile("data/15.log")
	if err != nil {
		log.Fatal(err.Error())
	}

	s2 := m2[_15.Coordinate{}]
	e2 := m2[_15.Coordinate{X: l2 - 1, Y: l2 - 1}]

	p2, err := g2.Shortest(s2, e2)

	fmt.Printf("Part 2 : %d\n", p2.Distance)
}
