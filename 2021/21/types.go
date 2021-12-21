package _21

type Player struct {
	Score, Position int
}

func (p *Player) Forward(n int) {
	p.Position = (p.Position + n) % 10
	if p.Position == 0 {
		p.Position = 10
	}
	p.Score += p.Position
}

func (p Player) Copy() Player {
	return Player{Position: p.Position, Score: p.Score}
}

type Die interface {
	Roll() int
}

type DeterministicDie struct {
	nextValue int
}

func NewDeterministicDie() *DeterministicDie {
	return &DeterministicDie{nextValue: 1}
}

func (d *DeterministicDie) Roll() int {
	r := d.nextValue

	d.nextValue++
	if d.nextValue > 100 {
		d.nextValue = 1
	}

	return r
}

type State struct {
	P1, P2 Player
}

type StateResult struct {
	Win1, Win2 int
}
