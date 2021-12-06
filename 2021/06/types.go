package _6

type SchoolOfFish []int

func InitSchoolOfFish(timers []int) SchoolOfFish {
	m := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := range timers {
		m[timers[i]]++
	}

	return m
}

func (s *SchoolOfFish) LiveAnotherDay() {
	newborns := (*s)[0]

	var tmp []int

	tmp = (*s)[1:9]
	tmp = append(tmp, newborns)
	tmp[6] += newborns
	*s = tmp
}

func (s *SchoolOfFish) LiveFor(n int) {
	for i := 0; i < n; i++ {
		s.LiveAnotherDay()
	}
}

func (s *SchoolOfFish) Population() int {
	a := 0

	for i := 0; i < 9; i++ {
		a += (*s)[i]
	}

	return a
}
