package _18

import (
	"log"
	"strconv"
	"strings"
)

func Add(s1, s2 string) string {
	sr := "[" + s1 + "," + s2 + "]"

	chars := strings.Split(sr, "")

	for {
		explosionIndex, splitIndex := DetectActions(chars)

		if explosionIndex != -1 {
			chars = Explode(chars, explosionIndex)
			sr = strings.Join(chars, "")
			continue
		}

		if splitIndex != -1 {
			chars = Split(chars, splitIndex)
			sr = strings.Join(chars, "")
			continue
		}

		return sr
	}
}

func DetectActions(chars []string) (int, int) {
	explosionStartIndex := -1
	splitIndex := -1
	nestSize := 0

	for i, c := range chars {
		if c == "[" {
			nestSize++
			if nestSize > 4 {
				explosionStartIndex = i
				continue
			}
		}

		if c == "]" {
			nestSize--
			if nestSize == 4 {
				break
			}
		}

		if c != "," {
			if len(c) > 1 && splitIndex == -1 {
				splitIndex = i
			}
		}
	}

	return explosionStartIndex, splitIndex
}

func Split(chars []string, index int) []string {
	n, err := strconv.Atoi(chars[index])
	if err != nil {
		panic(err.Error())
	}

	l := n / 2
	r := l + n%2

	tmp := make([]string, 0)
	tmp = append(tmp, chars[0:index]...)
	tmp = append(tmp, "[", strconv.Itoa(l), ",", strconv.Itoa(r), "]")
	tmp = append(tmp, chars[index+1:]...)

	return tmp
}

func Explode(chars []string, start int) []string {
	l, err := strconv.Atoi(chars[start+1])
	if err != nil {
		panic(err.Error())
	}

	r, err := strconv.Atoi(chars[start+3])
	if err != nil {
		panic(err.Error())
	}

	for i := start - 1; i > 0; i-- {
		if chars[i] != "[" && chars[i] != "]" && chars[i] != "," {
			n, err := strconv.Atoi(chars[i])
			if err != nil {
				panic(err.Error())
			}

			chars[i] = strconv.Itoa(n + l)
			break
		}
	}

	for i := start + 5; i < len(chars); i++ {
		if chars[i] != "[" && chars[i] != "]" && chars[i] != "," {
			n, err := strconv.Atoi(chars[i])
			if err != nil {
				panic(err.Error())
			}

			chars[i] = strconv.Itoa(n + r)
			break
		}
	}

	tmp := make([]string, 0)
	tmp = append(tmp, chars[0:start]...)
	tmp = append(tmp, "0")
	tmp = append(tmp, chars[start+5:]...)

	return tmp
}

func TransformToTreeAndReturnMagnitude(sfString string) int {
	chars := strings.Split(sfString, "")

	root := &Node{}

	var err error

	for _, c := range chars {
		if c == "[" {
			root.Left = &Node{Parent: root}
			root = root.Left
			continue
		}

		if c == "]" {
			m := root.Magnitude
			root = root.Parent
			root.Magnitude += 2 * m
			continue
		}

		if c == "," {
			m := root.Magnitude
			root = root.Parent
			root.Magnitude += 3 * m
			root.Right = &Node{Parent: root}
			root = root.Right
			continue
		}

		root.Magnitude, err = strconv.Atoi(c)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	return root.Magnitude
}

func LargestMagnitude(strings []string) int {
	max := 0

	for i := 0; i < len(strings)-1; i++ {
		for j := i + 1; j < len(strings); j++ {
			n := TransformToTreeAndReturnMagnitude(Add(strings[i], strings[j]))
			m := TransformToTreeAndReturnMagnitude(Add(strings[j], strings[i]))

			if n > m {
				if n > max {
					max = n
				}

				continue
			}

			if m > max {
				max = m
			}
		}
	}

	return max
}
