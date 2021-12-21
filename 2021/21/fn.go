package _21

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func NewPlayerFromPosition(fname string) (*Player, *Player, error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, nil, err
	}

	positions := strings.Split(strings.Replace(strings.Replace(string(b), "Player 1 starting position: ", "", 1), "Player 2 starting position: ", "", 1), "\n")
	p1, err := strconv.Atoi(positions[0])
	if err != nil {
		return nil, nil, err
	}

	p2, err := strconv.Atoi(positions[1])
	if err != nil {
		return nil, nil, err
	}

	return &Player{Position: p1}, &Player{Position: p2}, nil
}

func Play(p1, p2 *Player, d Die, scoreToWin int) (*Player, *Player, int) {
	turn := 0
	var winner, looser *Player

	for winner == nil {
		turn++
		if turn%2 == 1 {
			if playTurn(p1, d, scoreToWin) {
				winner = p1
				looser = p2
			}
		} else if playTurn(p2, d, scoreToWin) {
			winner = p2
			looser = p1
		}
	}

	return winner, looser, turn
}

func playTurn(p *Player, d Die, scoreToWin int) bool {
	v := d.Roll() + d.Roll() + d.Roll()
	p.Forward(v)
	if p.Score >= scoreToWin {
		return true
	}

	return false
}

func PlayWithMultiverse(p1, p2 *Player) (int, int) {
	outcomes := make([]int, 0, 27)
	for d1 := 1; d1 < 4; d1++ {
		for d2 := 1; d2 < 4; d2++ {
			for d3 := 1; d3 < 4; d3++ {
				outcomes = append(outcomes, d1+d2+d3)
			}
		}
	}

	cache := map[State]StateResult{}
	return countWinGamesForEachPlayer(*p1, *p2, outcomes, cache)
}

func countWinGamesForEachPlayer(p1, p2 Player, possibleOutcomes []int, cache map[State]StateResult) (int, int) {
	if p1.Score >= 21 {
		return 1, 0
	}

	if p2.Score >= 21 {
		return 0, 1
	}

	if v, ok := cache[State{p1, p2}]; ok {
		return v.Win1, v.Win2
	}

	var winCount1, winCount2 int

	for _, outcome := range possibleOutcomes {
		p := p1.Copy()
		p.Forward(outcome)

		w2, w1 := countWinGamesForEachPlayer(p2, p, possibleOutcomes, cache)
		winCount1 += w1
		winCount2 += w2
	}

	cache[State{p1, p2}] = StateResult{Win1: winCount1, Win2: winCount2}
	return winCount1, winCount2
}
