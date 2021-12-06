package _4

type Board struct {
	Rows    [][]int
	Columns [][]int
	Sum     int
}

func (b *Board) Pop(n int) bool {
	done := false
	rows := make([][]int, 0, 5)
	columns := make([][]int, 0, 5)

	for i := range b.Rows {
		r, checked := check(b.Rows[i], n)
		if checked {
			done = len(r) == 0
			b.Sum -= n
		}
		rows = append(rows, r)
	}

	b.Rows = rows

	for i := range b.Columns {
		c, checked := check(b.Columns[i], n)
		if checked {
			done = len(c) == 0 || done
		}
		columns = append(columns, c)
	}

	b.Columns = columns

	return done
}

func check(s []int, n int) ([]int, bool) {
	for index := range s {
		if s[index] == n {
			newSlice := remove(s, index)
			return newSlice, true
		}
	}

	return s, false
}

func remove(a []int, i int) []int {
	return append(a[:i], a[i+1:]...)
}
