package _8

var SignalIDs = map[rune]int{
	'a': 2,
	'b': 3,
	'c': 5,
	'd': 7,
	'e': 11,
	'f': 13,
	'g': 17,
}

type Output int

type SevenSegmentDisplay struct {
	Signal  []Output
	Digit   []Output
	Mapping map[Output]rune
}

func (s *SevenSegmentDisplay) Decrypt() {
	mor := map[Output]rune{}
	mro := map[rune]Output{}

	// Loop through the digits (result) and fill the map.
	shortGuess(s, mor, mro)

	// If the map is already filled with 4 runes, that means all the result has been decrypted => Return Map.
	if len(mor) == 4 {
		s.Mapping = mor
		return
	}

	bigGuess(s, mor, mro)

	s.Mapping = mor
}

/**
 * We're using the guessed runes to guess the next ones.
 * So we're using the starters (1,4,7,8) to guess the others, particularly 4 & 7.
 *
 *      ╭- 4 - 9 - (+7)5 - 6
 * [*] -|- ( 1, 8, 0, 2)
 *      ╰- 7 - 3
 *
 * We're using 4 to guess 9, then 9 and 7 to guess 5 and 5 to guess 6.
 */
func bigGuess(s *SevenSegmentDisplay, mor map[Output]rune, mro map[rune]Output) {
	i := 0
	for len(mor) != 10 {
		i = i % 10
		l2 := len(mor)
		_, ok := mor[s.Signal[i]]
		if !ok {
			l := len([]rune(s.Signal[i].String()))
			if l == 2 {
				mor[s.Signal[i]] = '1'
				mro['1'] = s.Signal[i]
				i++
				continue
			}

			if l == 3 {
				mor[s.Signal[i]] = '7'
				mro['7'] = s.Signal[i]
				i++
				continue
			}

			if l == 4 {
				mor[s.Signal[i]] = '4'
				mro['4'] = s.Signal[i]
				i++
				continue
			}

			if l == 5 {
				if o, f := mro['7']; f && s.Signal[i]%o == 0 {
					mor[s.Signal[i]] = '3'
					mro['3'] = s.Signal[i]
					i++
					continue
				}

				o, f := mro['9']
				o2, f2 := mro['7']
				if f && o%s.Signal[i] == 0 && f2 && s.Signal[i]%o2 != 0 {
					mor[s.Signal[i]] = '5'
					mro['5'] = s.Signal[i]
					i++
					continue
				}

				if l2 == 8 {
					mor[s.Signal[i]] = '2'
					i++
					continue
				}
			}

			if l == 6 {
				if o, f := mro['5']; f && s.Signal[i]%o == 0 {
					mor[s.Signal[i]] = '6'
					mro['6'] = s.Signal[i]
					i++
					continue
				}

				if o, f := mro['4']; f && s.Signal[i]%o == 0 {
					mor[s.Signal[i]] = '9'
					mro['9'] = s.Signal[i]
					i++
					continue
				}

				if l2 == 9 {
					mor[s.Signal[i]] = '0'
					i++
					continue
				}
			}

			if l == 7 {
				mor[s.Signal[i]] = '8'
				mro['8'] = s.Signal[i]
				i++
				continue
			}
		}
		i++
	}
}

func shortGuess(s *SevenSegmentDisplay, mor map[Output]rune, mro map[rune]Output) {
	for i := range s.Digit {
		l := len([]rune(s.Digit[i].String()))
		if l == 2 {
			mor[s.Digit[i]] = '1'
			mro['1'] = s.Digit[i]
			continue
		}

		if l == 3 {
			mor[s.Digit[i]] = '7'
			mro['7'] = s.Digit[i]
			continue
		}

		if l == 4 {
			mor[s.Digit[i]] = '4'
			mro['4'] = s.Digit[i]
			continue
		}

		if l == 7 {
			mor[s.Digit[i]] = '8'
			mro['8'] = s.Digit[i]
			continue
		}
	}
}

func (s *SevenSegmentDisplay) Read() string {
	str := ""

	for i := range s.Digit {
		str += string(s.Mapping[s.Digit[i]])
	}

	return str
}

func (s *SevenSegmentDisplay) CountDigitsWithUniqueSegmentsAmount() int {
	a := 0

	for i := range s.Digit {
		l := len([]rune(s.Digit[i].String()))

		if l == 2 || l == 3 || l == 4 || l == 7 {
			a++
		}
	}

	return a
}

func (o Output) String() string {
	s := ""

	if int(o)%SignalIDs['a'] == 0 {
		s += "a"
	}

	if int(o)%SignalIDs['b'] == 0 {
		s += "b"
	}

	if int(o)%SignalIDs['c'] == 0 {
		s += "c"
	}

	if int(o)%SignalIDs['d'] == 0 {
		s += "d"
	}

	if int(o)%SignalIDs['e'] == 0 {
		s += "e"
	}

	if int(o)%SignalIDs['f'] == 0 {
		s += "f"
	}

	if int(o)%SignalIDs['g'] == 0 {
		s += "g"
	}

	return s
}
