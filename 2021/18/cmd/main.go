package main

import (
	_18 "2021/18"
	"2021/common/reader"
	"fmt"
	"log"
)

func main() {
	strings, err := reader.ReadStringFileAndSplitByLine("data/18.log")
	if err != nil {
		log.Fatal(err.Error())
	}

	r := strings[0]
	for i := 1; i < len(strings); i++ {
		r = _18.Add(r, strings[i])
	}

	fmt.Printf("Part 1 : %d\n", _18.TransformToTreeAndReturnMagnitude(r))
	fmt.Printf("Part 2 : %d\n", _18.LargestMagnitude(strings))

}
