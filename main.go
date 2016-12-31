package main

import (
	"math/rand"

	tl "github.com/JoelOtter/termloop"
	"github.com/andoco/gotermlife/sim"
)

type SimLevel struct {
	*tl.BaseLevel
	sim *sim.S
}

var liveCell *tl.Cell

func (sl *SimLevel) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeySpace:
			sl.sim.Tick()
		}
	}
}

func (sl *SimLevel) Draw(s *tl.Screen) {
	for _, c := range sl.sim.Cells {
		if c.Live {
			s.RenderCell(c.Pos.X, c.Pos.Y, liveCell)
		}
	}
}

func main() {
	liveCell = &tl.Cell{Ch: '◼'}

	s := sim.New()
	//s.Seed([]sim.P{{5, 5}, {6, 5}, {7, 5}})

	seed := []sim.P{}
	for i := 0; i < 1000; i++ {
		seed = append(seed, sim.P{rand.Intn(80), rand.Intn(80)})
	}

	s.Seed(seed)

	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{})

	simLevel := &SimLevel{level, s}

	game.Screen().SetLevel(simLevel)
	game.Start()
}
