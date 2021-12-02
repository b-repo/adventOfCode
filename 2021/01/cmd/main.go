//Package main starts the program.
package main

import (
	"2021/01"
	"2021/common/reader"
	"fmt"
	"strconv"
)

func main() {
	data, err := reader.ReadIntFileAndSplitByLine("data/01.log")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("part 1 : " + strconv.Itoa(_1.IncreasedAmount(data)))
	fmt.Println("part 2 : " + strconv.Itoa(_1.IncreasedAmountThreeMeasurement(data)))
}
