package main

import (
	"math/rand"

	tl "github.com/JoelOtter/termloop"
	"github.com/andoco/gotermlife/sim"
)

type SimLevel struct {
	*tl.BaseLevel
	sim     *sim.S
	offsetX int
	offsetY int
}

var liveCell *tl.Cell

func (sl *SimLevel) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeySpace:
			sl.sim.Tick()
		case tl.KeyArrowUp:
			sl.offsetY += 1
		case tl.KeyArrowDown:
			sl.offsetY -= 1
		case tl.KeyArrowLeft:
			sl.offsetX += 1
		case tl.KeyArrowRight:
			sl.offsetX -= 1
		}
	}

	sl.sim.Tick()
}

func (sl *SimLevel) Draw(s *tl.Screen) {
	for _, c := range sl.sim.Cells {
		if c.Live {
			s.RenderCell(sl.offsetX+c.Pos.X, sl.offsetY+c.Pos.Y, liveCell)
		}
	}
}

func main() {
	liveCell = &tl.Cell{Ch: '◼'}

	s := sim.New()
	//s.Seed([]sim.P{{5, 5}, {6, 5}, {7, 5}})

	seed := []sim.P{}
	for i := 0; i < 1000; i++ {
		seed = append(seed, sim.P{rand.Intn(80), rand.Intn(40)})
	}
	s.Seed(seed)

	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{})

	simLevel := &SimLevel{level, s, 0, 0}

	game.Screen().SetLevel(simLevel)
	game.Screen().SetFps(10)
	game.Start()
}
