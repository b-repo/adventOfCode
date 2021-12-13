package main

import (
	_13 "2021/13"
	"fmt"
	"strconv"
)

func main() {
	t1, err := _13.ParseTransparentPaperFromFile("data/13.log")
	if err != nil {
		panic("unable to read file.")
	}

	t2, err := _13.ParseTransparentPaperFromFile("data/13.log")
	if err != nil {
		panic("unable to read file.")
	}

	i1 := t1.Instructions[0]
	t1.Process([]_13.Instruction{i1})

	t2.Process(t2.Instructions)

	fmt.Println("Part 1 : " + strconv.Itoa(t1.Count()))
	fmt.Println("Part 2 : ")
	t2.Display()
}
