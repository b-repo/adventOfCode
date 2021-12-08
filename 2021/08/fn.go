package _8

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadSignalsAndDigits(fname string) ([]SevenSegmentDisplay, error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	ssd := make([]SevenSegmentDisplay, 0, len(lines))

	for i := range lines {
		s := make([]Output, 0, len(lines))
		d := make([]Output, 0, 4)
		signalsAndDigits := strings.Split(lines[i], " | ")
		id := 1
		for _, c := range signalsAndDigits[0] + " " {
			v, found := SignalIDs[c]
			if !found {
				s = append(s, Output(id))
				id = 1
				continue
			}
			id *= v
		}

		for _, c := range signalsAndDigits[1] + " " {
			v, found := SignalIDs[c]
			if !found {
				d = append(d, Output(id))
				id = 1
				continue
			}
			id *= v
		}

		ssd = append(ssd, SevenSegmentDisplay{
			Signal: s,
			Digit:  d,
		})
	}

	return ssd, nil
}

func CountDigitsWithUniqueSegmentsAmount(s []SevenSegmentDisplay) int {
	a := 0

	for i := range s {
		a += s[i].CountDigitsWithUniqueSegmentsAmount()
	}

	return a
}

func SumDrecryptedDigits(s []SevenSegmentDisplay) int {
	sum := 0

	for i := range s {
		s[i].Decrypt()
		n, err := strconv.Atoi(s[i].Read())
		if err != nil {
			panic("unable to convert to int")
		}

		sum += n
	}

	return sum
}
