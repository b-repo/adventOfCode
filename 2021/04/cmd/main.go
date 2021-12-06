package main

import (
	_4 "2021/04"
	"2021/common/reader"
	"fmt"
	"strconv"
)

func main() {
	draw := []int{18, 99, 39, 89, 0, 40, 52, 72, 61, 77, 69, 51, 30, 83, 20, 65, 93, 88, 29, 22, 14, 82, 53, 41, 76, 79, 46, 78, 56, 57, 24, 36, 38, 11, 50, 1, 19, 26, 70, 4, 54, 3, 84, 33, 15, 21, 9, 58, 64, 85, 10, 66, 17, 43, 31, 27, 2, 5, 95, 96, 16, 97, 12, 34, 74, 67, 86, 23, 49, 8, 59, 45, 68, 91, 25, 48, 13, 28, 81, 94, 92, 42, 7, 37, 75, 32, 6, 60, 63, 35, 62, 98, 90, 47, 87, 73, 44, 71, 55, 80}
	data, err := reader.ReadStringFileAndSplitByLine("data/04.log")
	if err != nil {
		panic(err.Error())
	}

	b := _4.ParseBoards(data)
	s1 := _4.GetWinnerBoardID(b, draw)
	s2 := _4.GetLooserBoardID(b, draw)

	fmt.Println("Part 1 : " + strconv.Itoa(s1))
	fmt.Println("Part 2 : " + strconv.Itoa(s2))
}
