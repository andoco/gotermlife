package sim

import "testing"

func TestNew(t *testing.T) {
	sim := New()

	if sim == nil {
		t.Fatalf("expected S instance, got nil")
	}

	if len(sim.Cells) != 0 {
		t.Errorf("expected no cells, got %d", len(sim.Cells))
	}
}

func TestSeed(t *testing.T) {
	seedCells := []P{{0, 0}, {1, 0}}
	sim := New()
	sim.Seed(seedCells)

	if len(sim.Cells) != len(seedCells) {
		t.Errorf("expected %d cells, got %d", len(seedCells), len(sim.Cells))
	}

	for _, c := range sim.Cells {
		if !c.Live {
			t.Errorf("expected cell %v to be live, but was dead", c.Pos)
		}
	}
}

func TestNeighbours(t *testing.T) {
	testCases := []struct {
		pos        P
		neighbours []P
		name       string
	}{
		{
			P{0, 0},
			[]P{P{0, -1}, P{1, -1}, P{1, 0}, P{1, 1}, P{0, 1}, P{-1, 1}, P{-1, 0}, P{-1, -1}},
			"origin",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			n := neighbours(tc.pos)
			if len(n) != len(tc.neighbours) {
				t.Fatalf("expected %d neighbours, got %d", len(tc.neighbours), len(n))
			}

			for i, _ := range tc.neighbours {
				if tc.neighbours[i] != n[i] {
					t.Errorf("expected %d, got %d", tc.neighbours[i], n[i])
				}
			}
		})
	}
}

func TestBuildNeighbourCounts(t *testing.T) {
	cells := make(map[P]*C)
	cells[P{0, 0}] = &C{P{0, 0}, true}
	counts := buildNeighbourCounts(cells)

	if len(counts) != 8 {
		t.Fatalf("expected 8 counts, got %d", len(counts))
	}

	for _, v := range counts {
		if v != 1 {
			t.Fatalf("expected count of 1, got %d", v)
		}
	}
}

func TestApplyRules(t *testing.T) {
	testCases := []struct {
		neighboursAlive int
		inAlive         bool
		outAlive        bool
		name            string
	}{
		{
			1,
			true,
			false,
			"cell with 1 live neighbour dies",
		},
		{
			0,
			true,
			false,
			"cell with 0 live neighbours dies",
		},
		{
			2,
			true,
			true,
			"cell with 2 live neighbours lives",
		},
		{
			3,
			true,
			true,
			"cell with 3 live neighbours lives",
		},
		{
			4,
			true,
			false,
			"cell with more than 3 live neighbours dies",
		},
		{
			3,
			false,
			true,
			"dead cell with 3 live neighbours becomes alive",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			alive := applyRules(tc.inAlive, tc.neighboursAlive)
			if alive != tc.outAlive {
				t.Errorf("expected alive=%v, got %v", tc.outAlive, alive)
			}
		})
	}
}

func TestEnsureCell(t *testing.T) {
	m := make(map[P]*C)
	p := P{1, 1}
	ensureCell(m, p)
	c, ok := m[p]

	if !ok {
		t.Fatalf("expected cell, got none")
	}

	if c.Pos != p {
		t.Errorf("expected %v, got %v", p, c.Pos)
	}

	if c.Live {
		t.Error("expected dead cell, but was alive")
	}
}
