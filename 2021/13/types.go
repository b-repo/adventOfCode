package _13

import "fmt"

type Coordinate struct {
	X, Y int
}

type TransparentPaper struct {
	Grid         [][]bool
	Instructions []Instruction
}

type Instruction struct {
	IsHorizontalFold bool
	Value            int
}

func (t *TransparentPaper) GetInstructions(f func(Instruction) bool) []Instruction {
	instructions := make([]Instruction, 0, len(t.Instructions))

	for i := range t.Instructions {
		if f(t.Instructions[i]) {
			instructions = append(instructions, t.Instructions[i])
		}
	}

	return instructions
}

func (t *TransparentPaper) Process(instructions []Instruction) {
	for i := range instructions {
		if instructions[i].IsHorizontalFold {
			t.FoldHorizontally(instructions[i].Value)

			continue
		}
		t.FoldVertically(instructions[i].Value)
	}
}

func (t *TransparentPaper) Count() int {
	a := 0

	for i := range t.Grid {
		for j := range t.Grid[i] {
			if t.Grid[i][j] {
				a++
			}
		}
	}

	return a
}

func (t *TransparentPaper) FoldHorizontally(value int) {
	for i := 1; i <= value; i++ {
		if value-i < 0 || value+i >= len(t.Grid) {
			break
		}

		for x := range t.Grid[value-i] {
			t.Grid[value-i][x] = t.Grid[value+i][x] || t.Grid[value-i][x]
		}
	}

	t.Grid = t.Grid[:value]
}

func (t *TransparentPaper) FoldVertically(value int) {
	for i := range t.Grid {
		for x := 1; x <= value; x++ {
			if value-x < 0 || value+x >= len(t.Grid[i]) {
				break
			}

			t.Grid[i][value-x] = t.Grid[i][value+x] || t.Grid[i][value-x]
		}

		t.Grid[i] = t.Grid[i][:value]
	}
}

func (t *TransparentPaper) Display() {
	for y := range t.Grid {
		for x := range t.Grid[y] {
			if t.Grid[y][x] {
				fmt.Print("#")

				continue
			}
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}
