package main

import (
	_19 "2021/19"
	"fmt"
	"log"
)

func main() {
	scanners, vectors, err := _19.ParseScannerFromFile("data/19.log")
	if err != nil {
		log.Fatal(err.Error())
	}

	uniqueBeacons, finalPositions := _19.PlaceScanners(scanners, vectors)

	fmt.Printf("Part 1 : %d\n", uniqueBeacons)
	fmt.Printf("Part 2 : %d\n", _19.CalculateMaxDistance(scanners, finalPositions))
}
