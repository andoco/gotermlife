package sim

func New() *S {
	return &S{}
}

// S holds the state of the simulated universe.
type S struct {
	Cells []C
}

// Seed sets the initial live cells in the simulation.
func (s *S) Seed(cells []P) {
	for _, p := range cells {
		s.Cells = append(s.Cells, C{p, true})
	}
}

// Tick will calculate the next iteration of the simulation universe, and assign to C.
func (s *S) Tick() {
}

// C is the state of a cell in the simulation universe.
type C struct {
	Pos  P
	Live bool
}

// P is a position in the universe.
type P struct {
	X, Y int
}

func tickCells(cells []C) []C {
	return []C{}
}
