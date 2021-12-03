package main

import (
	_3 "2021/03"
	"2021/common/reader"
	"fmt"
	"strconv"
)

func main() {
	data, err := reader.Read2DStringSliceFile("data/03.log")
	if err != nil {
		panic(err.Error())
	}

	energy := _3.ReadEnergyConsumption(data)
	lifeSupport := _3.ReadLifeSupportRating(data)

	fmt.Println("Part 1 : " + strconv.FormatInt(energy.GammaRate*energy.EpsilonRate, 10))
	fmt.Println("Part 2 : " + strconv.FormatInt(lifeSupport.O2GeneratorRate*lifeSupport.CO2ScrubberRate, 10))
}
