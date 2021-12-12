package _12

import (
	"2021/common/reader"
	"strings"
)

func ReadCavesMapFromFile(fname string) (map[string][]string, error) {
	d, err := reader.ReadStringFileAndSplitByLine(fname)
	if err != nil {
		return nil, err
	}

	m := map[string][]string{}

	for i := range d {
		tunnel := strings.Split(d[i], "-")

		if tunnel[0] == "start" || tunnel[1] == "end" {
			m[tunnel[0]] = append(m[tunnel[0]], tunnel[1])

			continue
		}

		if tunnel[1] == "start" || tunnel[0] == "end" {
			m[tunnel[1]] = append(m[tunnel[1]], tunnel[0])

			continue
		}

		m[tunnel[0]] = append(m[tunnel[0]], tunnel[1])
		m[tunnel[1]] = append(m[tunnel[1]], tunnel[0])
	}

	return m, nil
}

func Explore(m map[string][]string, cave, path string) []string {
	possiblePaths, f := m[cave]
	if !f {
		return []string{path}
	}

	r := make([]string, 0)

	for i := range possiblePaths {
		if strings.ToLower(possiblePaths[i]) == possiblePaths[i] && strings.Contains(path, possiblePaths[i]) {
			continue
		}

		p := path + "," + possiblePaths[i]
		r = append(r, Explore(m, possiblePaths[i], p)...)
	}

	return r
}

func ExploreTwice(m map[string][]string, cave, path string) []string {
	possiblePaths, f := m[cave]
	if !f {
		return []string{path}
	}

	r := make([]string, 0)

	for i := range possiblePaths {
		if strings.ToLower(possiblePaths[i]) == possiblePaths[i] &&
			hasAlreadyVisitedASmallCaveTwice(possiblePaths[i], path) {
			continue
		}

		p := path + "," + possiblePaths[i]
		r = append(r, ExploreTwice(m, possiblePaths[i], p)...)
	}

	return r
}

func hasAlreadyVisitedASmallCaveTwice(e, path string) bool {
	occurrences := strings.Count(path, ","+e)

	if occurrences >= 2 {
		return true
	}

	if occurrences == 1 {
		m := map[string]int{}

		caves := strings.Split(path, ",")

		for i := range caves {
			m[caves[i]]++
			if strings.ToLower(caves[i]) == caves[i] && m[caves[i]] >= 2 {
				return true
			}
		}
	}

	return false
}
