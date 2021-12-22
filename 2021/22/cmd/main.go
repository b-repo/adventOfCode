package main

import (
	_22 "2021/22"
	"fmt"
	"log"
)

func main() {
	var core1, core2 _22.Core
	var err error

	core1, err = _22.ParseCore("data/22.log", func(cuboid _22.Cuboid) bool {
		return cuboid.Min.X >= -50 && cuboid.Min.Y >= -50 && cuboid.Min.Z >= -50 &&
			cuboid.Max.X <= 50 && cuboid.Max.Y <= 50 && cuboid.Max.Z <= 50
	})
	if err != nil {
		log.Fatal(err)
	}

	core2, err = _22.ParseCore("data/22.log", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 : %d\n", core1.Compute())
	fmt.Printf("Part 2 : %d\n", core2.Compute())

}
