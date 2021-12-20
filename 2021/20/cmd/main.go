package main

import (
	_20 "2021/20"
	"fmt"
	"log"
)

func main() {
	a, i, err := _20.ParseInput("data/20.log")
	if err != nil {
		log.Fatal(err.Error())
	}

	_20.EnhanceImage(2, i, a)

	fmt.Printf("Part 1 : %d\n", i.LitPixels)

	_20.EnhanceImage(48, i, a)

	fmt.Printf("Part 2 : %d\n", i.LitPixels)

}
