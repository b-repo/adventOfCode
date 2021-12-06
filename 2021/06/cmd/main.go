package main

import (
	_6 "2021/06"
	"2021/common/reader"
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	data, err := reader.ReadIntFileAndSplitBy("data/06.log", ",")
	if err != nil {
		panic(err.Error())
	}

	s := _6.InitSchoolOfFish(data)
	s.LiveFor(80)

	s2 := _6.InitSchoolOfFish(data)
	s2.LiveFor(256)

	fmt.Println("Part 1 : " + strconv.Itoa(s.Population()))
	fmt.Println("Part 2 : " + strconv.Itoa(s2.Population()))
	fmt.Println(time.Since(start))
}
