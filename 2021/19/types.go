package _19

type Scanner []Beacon

type Beacon [48]Coordinate

type ScannerVectors []BeaconVector

type BeaconVector [48]VectorHash

type VectorHash []uint64

type Coordinate struct {
	X, Y, Z int
}

func (p Coordinate) VecTo(q Coordinate) Vector {
	return Vector{q.X - p.X, q.Y - p.Y, q.Z - p.Z}
}

func (p Coordinate) PossibleOrientations() [48]Coordinate {
	var res [48]Coordinate
	x, y, z := p.X, p.Y, p.Z
	for i := 0; i < 8; i++ {
		j := i * 6
		res[j] = Coordinate{x, y, z}
		res[j+1] = Coordinate{x, z, y}
		res[j+2] = Coordinate{y, x, z}
		res[j+3] = Coordinate{y, z, x}
		res[j+4] = Coordinate{z, x, y}
		res[j+5] = Coordinate{z, y, x}
		if i&1 == 1 {
			for k := j; k < j+6; k++ {
				res[k].X = -res[k].X
			}
		}
		if i&2 == 2 {
			for k := j; k < j+6; k++ {
				res[k].Y = -res[k].Y
			}
		}
		if i&4 == 4 {
			for k := j; k < j+6; k++ {
				res[k].Z = -res[k].Z
			}
		}
	}
	return res
}

type Vector struct {
	X, Y, Z int
}

func (v Vector) Hash() uint64 {
	// Random prime numbers
	return uint64((v.X + 2551) + 4889*(v.Y+2551) + 4889*4721*(v.Z+2551))
}
