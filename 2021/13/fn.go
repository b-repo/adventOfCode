package _13

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ParseTransparentPaperFromFile(fname string) (*TransparentPaper, error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	d := strings.Split(string(b), "\n\n")

	var xMax, yMax int

	strC := strings.Split(d[0], "\n")
	coordinates := make([]Coordinate, 0, len(strC))

	for i := range strC {
		c := strings.Split(strC[i], ",")

		x, err := strconv.Atoi(c[0])
		if err != nil {
			return nil, err
		}

		y, err := strconv.Atoi(c[1])
		if err != nil {
			return nil, err
		}

		coordinates = append(coordinates, Coordinate{x, y})

		if x > xMax {
			xMax = x
		}

		if y > yMax {
			yMax = y
		}
	}

	board := make([][]bool, yMax+1)
	for i := range board {
		board[i] = make([]bool, xMax+1)
	}

	for i := range coordinates {
		board[coordinates[i].Y][coordinates[i].X] = true
	}

	strI := strings.Split(d[1], "\n")

	instructions := make([]Instruction, 0, len(strI))

	for i := range strI {
		tmp := strings.Replace(strI[i], "fold along ", "", 1)
		instruction := strings.Split(tmp, "=")

		v, err := strconv.Atoi(instruction[1])
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, Instruction{
			IsHorizontalFold: instruction[0] == "y",
			Value:            v,
		})
	}

	return &TransparentPaper{
		Grid:         board,
		Instructions: instructions,
	}, nil
}
