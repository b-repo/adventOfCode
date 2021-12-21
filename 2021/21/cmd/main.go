package main

import (
	_21 "2021/21"
	"fmt"
	"log"
)

func main() {
	p1, p2, err := _21.NewPlayerFromPosition("data/21.log")
	if err != nil {
		log.Fatal(err.Error())
	}

	_, looser, playedTurns := _21.Play(p1, p2, _21.NewDeterministicDie(), 1000)

	fmt.Printf("Part 1 : %d\n", looser.Score*(playedTurns*3))

	p1, p2, err = _21.NewPlayerFromPosition("data/21.log")
	if err != nil {
		log.Fatal(err.Error())
	}

	p1Wins, p2Wins := _21.PlayWithMultiverse(p1, p2)
	var w int

	if p1Wins > p2Wins {
		w = p1Wins
	} else {
		w = p2Wins
	}

	fmt.Printf("Part 2 : %d\n", w)

}
