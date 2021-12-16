package _16

func ParseBits(bs []byte) Bits {
	r := make([]int, len(bs)*8)
	for i, b := range bs {
		for j := 0; j < 8; j++ {
			r[i*8+j] = int(b >> uint(7-j) & 0x01)
		}
	}
	return r
}

func Operate(t TypeID, old int64, new int64) int64 {
	switch t {
	case Sum:
		if old == -1 {
			old = 0
		}

		return old + new
	case Product:
		if old == -1 {
			old = 1
		}

		return old * new
	case Min:
		if old == -1 {
			return new
		}

		if old > new {
			return new
		}

		return old
	case Max:
		if old == -1 {
			return new
		}

		if old > new {
			return old
		}

		return new
	case Greater:
		if old == -1 {
			return new
		}

		if old > new {
			return 1
		}

		return 0
	case Lesser:
		if old == -1 {
			return new
		}

		if old < new {
			return 1
		}

		return 0
	case Equal:
		if old == -1 {
			return new
		}

		if old == new {
			return 1
		}

		return 0
	}

	return -1
}
