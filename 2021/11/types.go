package _11

type Octopus struct {
	Energy, X, Y int
	HasFlashed   bool
}

func (o *Octopus) Gain(b [][]Octopus, f *int) {
	if !o.HasFlashed {
		o.Energy++
		if o.Energy > 9 {
			o.Flash(b, f)
		}
	}
}

func (o *Octopus) Flash(b [][]Octopus, f *int) {
	*f++
	o.HasFlashed = true
	o.Energy = 0
	o.Irradiate(b, f)
}

func (o *Octopus) Irradiate(b [][]Octopus, f *int) {
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if o.X+i >= 0 && o.Y+j >= 0 &&
				o.X+i < len(b) && o.Y+j < len(b[0]) &&
				b[o.X+i][o.Y+j] != *o && !b[o.X+i][o.Y+j].HasFlashed {
				b[o.X+i][o.Y+j].Gain(b, f)
			}
		}
	}
}
