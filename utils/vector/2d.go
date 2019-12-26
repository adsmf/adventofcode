package vector

type FPoint struct {
	x, y float64
}

func (p FPoint) Add(a FPoint) FPoint {
	return FPoint{
		x: p.x + a.x,
		y: p.y + a.y,
	}
}

type GridPoint struct {
	X, Y float64
}

func (p GridPoint) Add(a GridPoint) GridPoint {
	return GridPoint{
		X: p.X + a.X,
		Y: p.Y + a.Y,
	}
}

func (p GridPoint) Up() GridPoint {
	return p.Add(GridPoint{0, 1})
}

func (p GridPoint) Down() GridPoint {
	return p.Add(GridPoint{0, -1})
}

func (p GridPoint) Left() GridPoint {
	return p.Add(GridPoint{-1, 0})
}

func (p GridPoint) Right() GridPoint {
	return p.Add(GridPoint{1, 0})
}
