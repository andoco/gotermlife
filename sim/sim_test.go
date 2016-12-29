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
