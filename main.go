package main

import "github.com/andoco/gotermlife/sim"
import tl "github.com/JoelOtter/termloop"

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
	liveCell = &tl.Cell{Fg: tl.ColorRed, Ch: '*'}

	s := sim.New()
	s.Seed([]sim.P{{5, 5}, {10, 10}})

	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: ' ',
	})

	simLevel := &SimLevel{level, s}

	game.Screen().SetLevel(simLevel)
	game.Start()
}
