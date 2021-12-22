package _19

import (
	"2021/common/reader"
	"math"
	"sort"
	"strconv"
	"strings"
)

func ParseScannerFromFile(fname string) ([]Scanner, []ScannerVectors, error) {
	s, err := reader.ReadStringFileAndSplitByLine(fname)
	if err != nil {
		return nil, nil, err
	}

	scanners := make([]Scanner, 0)

	var i, j int

	for i < len(s) {
		i++
		scanners = append(scanners, make(Scanner, 0, 28))

		for i < len(s) && s[i] != "" {
			coordinates := strings.Split(s[i], ",")
			x, err := strconv.Atoi(coordinates[0])
			if err != nil {
				return nil, nil, err
			}

			y, err := strconv.Atoi(coordinates[1])
			if err != nil {
				return nil, nil, err
			}

			z, err := strconv.Atoi(coordinates[2])
			if err != nil {
				return nil, nil, err
			}

			c := Coordinate{x, y, z}
			scanners[j] = append(scanners[j], c.PossibleOrientations())
			i++
		}

		i++
		j++
	}
	return scanners, CreateVectorHashes(scanners), nil
}

func CreateVectorHashes(scanners []Scanner) []ScannerVectors {
	scannerVectors := make([]ScannerVectors, len(scanners))

	minVectors := math.MaxInt32
	maxVectors := 0

	for scanner, beacons := range scanners {
		beaconsAmount := len(scanners[scanner])
		scannerVectors[scanner] = make(ScannerVectors, beaconsAmount)

		for beacon := 0; beacon < beaconsAmount; beacon++ {
			for orientation := 0; orientation < 48; orientation++ {
				scannerVectors[scanner][beacon][orientation] = make(VectorHash, 0, 27)
			}
		}

		minVectors = min(minVectors, beaconsAmount-1)
		maxVectors = max(maxVectors, beaconsAmount-1)
		for beacon1 := 0; beacon1 < beaconsAmount-1; beacon1++ {
			for beacon2 := beacon1 + 1; beacon2 < beaconsAmount; beacon2++ {
				b1, b2 := beacons[beacon1], beacons[beacon2]
				for orientation := range b1 {
					v12 := b1[orientation].VecTo(b2[orientation])
					scannerVectors[scanner][beacon1][orientation] = append(scannerVectors[scanner][beacon1][orientation], v12.Hash())
					v21 := b2[orientation].VecTo(b1[orientation])
					scannerVectors[scanner][beacon2][orientation] = append(scannerVectors[scanner][beacon2][orientation], v21.Hash())
				}
			}
		}

		for beacon := 0; beacon < beaconsAmount; beacon++ {
			for orientation := 0; orientation < 48; orientation++ {
				sort.Slice(scannerVectors[scanner][beacon][orientation], func(i, j int) bool {
					return scannerVectors[scanner][beacon][orientation][i] < scannerVectors[scanner][beacon][orientation][j]
				})
			}
		}
	}

	return scannerVectors
}

func CalculateMaxDistance(scanners []Scanner, finalPositions []Coordinate) int {
	var maxDist int

	for i := 0; i < len(scanners)-1; i++ {
		for j := i + 1; j < len(scanners); j++ {
			c1, c2 := finalPositions[i], finalPositions[j]
			maxDist = max(maxDist, abs(c2.X-c1.X)+abs(c2.Y-c1.Y)+abs(c2.Z-c1.Z))
		}
	}
	return maxDist
}

func PlaceScanners(scanners []Scanner, vectors []ScannerVectors) (int, []Coordinate) {
	scannerAmount := len(scanners)

	found := make([]bool, scannerAmount)
	found[0] = true
	counter := 1
	cur := []int{0}
	next := []int{}

	finalPositions := make([]Coordinate, scannerAmount)
	finalPositions[0] = Coordinate{0, 0, 0}

	for counter != scannerAmount {
		next = next[:0]
		for _, currentIndex := range cur {
			for otherIndex := 0; otherIndex < scannerAmount; otherIndex++ {
				if found[otherIndex] {
					continue
				}

				if v, ok := compareScanners(scanners, vectors, currentIndex, otherIndex); ok {
					finalPositions[otherIndex] = v
					found[otherIndex] = true
					next = append(next, otherIndex)
					counter++
				}
			}
		}

		cur, next = next, cur
	}

	uniquePoints := make(map[Coordinate]struct{})
	for scanner := range scanners {
		for beacon := range scanners[scanner] {
			uniquePoints[scanners[scanner][beacon][0]] = struct{}{}
		}
	}

	return len(uniquePoints), finalPositions
}

func compareScanners(scanners []Scanner, vectors []ScannerVectors, s1Index, s2Index int) (Coordinate, bool) {
	for b1 := range vectors[s1Index] {
		v := vectors[s1Index][b1][0]
		for b2 := range vectors[s2Index] {
			for orientation := range vectors[s2Index][b2] {
				if sharesSpace(v, vectors[s2Index][b2][orientation]) {
					c1 := scanners[s1Index][b1][0]
					c2 := scanners[s2Index][b2][orientation]
					dx, dy, dz := c2.X-c1.X, c2.Y-c1.Y, c2.Z-c1.Z

					for b2bis := range vectors[s2Index] {
						vectors[s2Index][b2bis][0] = vectors[s2Index][b2bis][orientation]
						scanners[s2Index][b2bis][0] = scanners[s2Index][b2bis][orientation]

						c := scanners[s2Index][b2bis][0]
						scanners[s2Index][b2bis][0] = Coordinate{
							X: c.X - dx,
							Y: c.Y - dy,
							Z: c.Z - dz,
						}
					}

					return Coordinate{-dx, -dy, -dz}, true
				}
			}
		}
	}
	return Coordinate{}, false
}

func sharesSpace(v1, v2 VectorHash) bool {
	var v1Index, v2Index, count int
	for v1Index != len(v1) && v2Index != len(v2) && min(len(v1), len(v2))+count >= 11 {
		if v1[v1Index] == v2[v2Index] {
			count++
			if count == 11 {
				return true
			}

			v1Index++
			v2Index++

			continue
		}

		if v1[v1Index] < v2[v2Index] {
			v1Index++

			continue
		}

		if v2[v2Index] < v1[v1Index] {
			v2Index++
		}
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
