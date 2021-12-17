package _17

type Vector struct {
	X, Y int
}

type Target struct {
	Xmin, Xmax, Ymin, Ymax int
}

func (v *Vector) WillReachTarget(t Target) bool {
	pX := 0
	pY := 0
	x := v.X
	y := v.Y
	for pY >= t.Ymin {
		pX += x
		if x > 0 {
			x--
		}

		pY += y
		y--

		if pX >= t.Xmin && pX <= t.Xmax && pY >= t.Ymin && pY <= t.Ymax {
			return true
		}
	}

	return false
}
