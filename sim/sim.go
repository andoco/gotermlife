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
		s.Cells[p] = &C{p, true}
	}
}

// Tick will calculate the next iteration of the simulation universe, and assign to C.
func (s *S) Tick() {
	neighbourCounts := make(map[P]int)

	for _, c := range s.Cells {
		if !c.Live {
			continue
		}

		for _, np := range neighbours(c.Pos) {
			neighbourCounts[np] += 1
			ensureCell(s.Cells, np)
		}
	}

	for _, c := range s.Cells {
		c.Live = applyRules(c.Live, neighbourCounts[c.Pos])

		if !c.Live {
			delete(s.Cells, c.Pos)
		}
	}
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

func neighbours(pos P) []P {
	n := []P{
		P{pos.X, pos.Y - 1},
		P{pos.X + 1, pos.Y - 1},
		P{pos.X + 1, pos.Y},
		P{pos.X + 1, pos.Y + 1},
		P{pos.X, pos.Y + 1},
		P{pos.X - 1, pos.Y + 1},
		P{pos.X - 1, pos.Y},
		P{pos.X - 1, pos.Y - 1},
	}

	return n
}

func ensureCell(m map[P]*C, p P) {
	_, ok := m[p]
	if !ok {
		m[p] = &C{p, false}
	}
}
