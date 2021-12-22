package _22

type Core []Cuboid

func (c Core) Count() int64 {
	var v int64

	for _, cuboid := range c {
		if cuboid.isOn {
			v += cuboid.Volume()
			continue
		}

		v -= cuboid.Volume()
	}

	return v
}

func (c Core) Compute() int64 {
	visited := make(Core, 0)

	for _, cuboid := range c {
		tmp := make(Core, 0)
		if cuboid.isOn {
			tmp = append(tmp, cuboid)
		}

		for _, cuboid1 := range visited {
			if v, ok := cuboid.Intersect(cuboid1); ok {
				tmp = append(tmp, v)
			}
		}

		visited = append(visited, tmp...)
	}

	return visited.Count()
}

type Cuboid struct {
	isOn     bool
	Min, Max Coordinate
}

func NewCuboid(xMin, xMax, yMin, yMax, zMin, zMax int64, state bool) Cuboid {
	return Cuboid{state, Coordinate{xMin, yMin, zMin}, Coordinate{xMax, yMax, zMax}}
}

func (c Cuboid) Intersect(c1 Cuboid) (Cuboid, bool) {
	r := Cuboid{
		isOn: !c1.isOn,
		Min:  Coordinate{max(c.Min.X, c1.Min.X), max(c.Min.Y, c1.Min.Y), max(c.Min.Z, c1.Min.Z)},
		Max:  Coordinate{min(c.Max.X, c1.Max.X), min(c.Max.Y, c1.Max.Y), min(c.Max.Z, c1.Max.Z)},
	}
	return r, r.Min.X <= r.Max.X && r.Min.Y <= r.Max.Y && r.Min.Z <= r.Max.Z
}

func (c Cuboid) Volume() int64 {
	return (abs(c.Max.X-c.Min.X) + 1) * (abs(c.Max.Y-c.Min.Y) + 1) * (abs(c.Max.Z-c.Min.Z) + 1)
}

type Coordinate struct {
	X, Y, Z int64
}
