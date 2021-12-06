package _5

func ParseLines(d []string, f func(l Line) bool) ([]Line, int) {
	lines := make([]Line, 0, len(d))
	max := 0

	for i := range d {
		l, m := ParseLine(d[i])
		if f(l) {
			if m > max {
				max = m
			}

			lines = append(lines, l)
		}
	}

	return lines, max
}

func DrawLines(l []Line, f *Field) {
	for i := range l {
		l[i].Draw(f)
	}
}
