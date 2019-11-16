package gol

// Grid abstraction of the grid for game of life
type Grid struct {
	Arr []uint8
	W   int
}

// IsValid check whether the Grid is a valid one
func (g *Grid) IsValid() bool {
	return len(g.Arr)%g.W == 0
}

// Next returns the next generation of the grid
func (g *Grid) Next() *Grid {
	return nil
}

// String returns display string of a Grid
func (g *Grid) String() string {
	s := ""
	for i, c := range g.Arr {
		if c == 1 {
			s += "*"
		} else if c == 0 {
			s += "."
		}
		if i > 0 && (i+1)%g.W == 0 {
			s += "\n"
		}
	}

	return s
}
