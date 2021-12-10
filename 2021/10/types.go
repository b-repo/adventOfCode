package _10

type Element struct {
	Character string
	Previous  *Element
}

func (e *Element) Opposite() string {
	switch e.Character {
	case "(":
		return ")"
	case "[":
		return "]"
	case "{":
		return "}"
	case "<":
		return ">"
	default:
		panic("unknown character")
	}
}

func (e *Element) IsOpener() bool {
	return e.Character == "(" || e.Character == "[" || e.Character == "{" || e.Character == "<"
}
