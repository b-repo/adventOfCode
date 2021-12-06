package _4

import (
	"strconv"
	"strings"
)

func ParseBoards(d []string) []Board {
	boards := make([]Board, 0)

	for start := 0; start < len(d)-4; start += 6 {
		boards = append(boards, parseBoard(d, start, start+5))
	}

	return boards
}

func parseBoard(data []string, start, end int) Board {
	columns := initColumns()
	rows := make([][]int, 0, 5)
	sum := 0

	for index := start; index < end; index++ {
		d := strings.Split(data[index], " ")
		r := make([]int, 0, 5)

		columnIndex := 0

		for i := range d {
			if d[i] == "" {
				continue
			}
			n, err := strconv.Atoi(d[i])
			if err != nil {
				panic("unable to convert to integer")
			}
			r = append(r, n)
			columns[columnIndex] = append(columns[columnIndex], n)
			sum += n
			columnIndex++
		}

		rows = append(rows, r)
	}

	return Board{
		Rows:    rows,
		Columns: columns,
		Sum:     sum,
	}
}

func initColumns() [][]int {
	c := make([][]int, 0, 5)
	for i := 0; i < 5; i++ {
		d := make([]int, 0, 5)
		c = append(c, d)
	}
	return c
}

func GetWinnerBoardID(boards []Board, draw []int) int {
	for _, d := range draw {
		for i := range boards {
			if boards[i].Pop(d) {
				return boards[i].Sum * d
			}
		}
	}

	return 0
}

func GetLooserBoardID(boards []Board, draw []int) int {
	for _, d := range draw {
		i := getNextBoardToPop(boards, d)
		if len(i) > 0 {
			if len(boards) == 1 {
				return boards[i[0]].Sum * d
			}

			shift := 0
			for index := range i {
				boards = append(boards[:i[index]-shift], boards[i[index]-shift+1:]...)
				shift++
			}
		}
	}
	return 0
}

func getNextBoardToPop(boards []Board, n int) []int {
	indexes := make([]int, 0)
	for i := range boards {
		if boards[i].Pop(n) {
			indexes = append(indexes, i)
		}
	}
	return indexes
}
