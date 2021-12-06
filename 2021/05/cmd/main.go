package main

import (
	_5 "2021/05"
	"2021/common/reader"
	"fmt"
	"strconv"
)

func main() {
	data, err := reader.ReadStringFileAndSplitByLine("data/05.log")
	if err != nil {
		panic(err.Error())
	}

	lines, fieldSize := _5.ParseLines(data, func(l _5.Line) bool {
		return l.Start.X == l.End.X || l.Start.Y == l.End.Y
	})
	field := _5.InitField(fieldSize)
	_5.DrawLines(lines, field)

	lines2, fieldSize := _5.ParseLines(data, func(l _5.Line) bool {
		return true
	})
	field2 := _5.InitField(fieldSize)
	_5.DrawLines(lines2, field2)

	fmt.Println("Part 1 : " + strconv.Itoa(field.CountOverlapOver(2)))
	fmt.Println("Part 2 : " + strconv.Itoa(field2.CountOverlapOver(2)))
}
