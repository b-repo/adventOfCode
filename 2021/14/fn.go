package _14

import (
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func ParsePolymerDataFromFile(fname string) (string, map[string]string) {
	f, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}

	d := strings.Split(string(f), "\n\n")

	tpl := d[0]

	m := map[string]string{}

	s := strings.Split(d[1], "\n")

	for i := range s {
		e := strings.Split(s[i], " -> ")
		m[e[0]] = e[1]
	}

	return tpl, m
}

func InsertPairs(tpl string, m map[string]string, n int) map[string]int {
	s := strings.Split(tpl, "")

	counter := map[string]int{}

	for i := 1; i < len(s); i++ {
		cpl := s[i-1] + s[i]
		_, ok := counter[cpl]
		if !ok {
			counter[cpl] = 1

			continue
		}

		counter[cpl]++
	}

	for i := 0; i < n; i++ {
		counter = InsertPair(counter, m)
	}

	return counter
}

func InsertPair(c map[string]int, m map[string]string) map[string]int {
	next := map[string]int{}

	for k := range c {
		_, ok := m[k]
		if ok {
			s := strings.Split(k, "")
			p1 := s[0] + m[k]
			p2 := m[k] + s[1]

			_, ok = next[p1]
			if !ok {
				next[p1] = c[k]
			} else {
				next[p1] += c[k]
			}

			_, ok = next[p2]
			if !ok {
				next[p2] = c[k]
			} else {
				next[p2] += c[k]
			}
		}
	}

	return next
}

func FindLeastAndMostCommonElementInPolymer(c map[string]int) (int, int) {
	min := math.MaxFloat64
	max := 0.
	counter := map[string]float64{}

	for k := range c {
		chars := strings.Split(k, "")
		v := float64(c[k]) / 2
		_, ok := counter[chars[0]]
		if !ok {
			counter[chars[0]] = v
		} else {
			counter[chars[0]] += v
		}

		_, ok = counter[chars[1]]
		if !ok {
			counter[chars[1]] = v
		} else {
			counter[chars[1]] += v
		}
	}

	for k := range counter {
		if counter[k] < min {
			min = counter[k]
		}

		if counter[k] > max {
			max = counter[k]
		}
	}

	return int(math.Round(min)), int(math.Round(max))
}
