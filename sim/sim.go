package sim

func New() *S {
	return &S{Cells: make(map[P]*C)}
}

// S holds the state of the simulated universe.
type S struct {
	Cells map[P]*C
}

// Seed sets the initial live cells in the simulation.
func (s *S) Seed(cells []P) {
	for _, p := range cells {
		s.Cells[p] = &C{p, 0, true}
	}
}

// Tick will calculate the next iteration of the simulation universe, and assign to C.
func (s *S) Tick() {
	for _, c := range s.Cells {
		c.NeighbourCount = 0
	}

	for _, c := range s.Cells {
		neighbours := s.neighbourCells(c.Pos)
		for _, nc := range neighbours {
			if c.Live {
				nc.NeighbourCount += 1
			}
		}
	}

	for _, c := range s.Cells {
		c.Live = applyRules(c.Live, c.NeighbourCount)
	}
}

// C is the state of a cell in the simulation universe.
type C struct {
	Pos            P
	NeighbourCount int
	Live           bool
}

// P is a position in the universe.
type P struct {
	X, Y int
}

func applyRules(live bool, neighbours int) bool {
	if live && neighbours < 2 {
		return false
	}

	if live && (neighbours == 2 || neighbours == 3) {
		return true
	}

	if live && neighbours > 3 {
		return false
	}

	if !live && neighbours == 3 {
		return true
	}

	return live
}

func (s *S) neighbourCells(pos P) []*C {
	get := func(pos P) *C {
		c, ok := s.Cells[pos]
		if !ok {
			c = &C{Pos: pos}
			s.Cells[pos] = c
		}
		return c
	}

	n := []*C{
		get(P{pos.X, pos.Y - 1}),
		get(P{pos.X + 1, pos.Y - 1}),
		get(P{pos.X + 1, pos.Y}),
		get(P{pos.X + 1, pos.Y + 1}),
		get(P{pos.X, pos.Y + 1}),
		get(P{pos.X - 1, pos.Y + 1}),
		get(P{pos.X - 1, pos.Y}),
		get(P{pos.X - 1, pos.Y - 1}),
	}

	return n
}
