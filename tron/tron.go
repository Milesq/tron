package tron

import (
	"time"

	"github.com/milesq/tron/tron/event"
)

// GameState .
type GameState struct {
	X int
}

// Game .
type Game struct {
	Exited bool
	State  GameState
}

// NewGame .
func NewGame(cfg Config) Game {
	return Game{}
}

// Emit .
func (tron *Game) Emit(ev event.Event) {
	switch ev {
	case event.Exit:
		tron.Exited = true
	}
}

// Next .
func (tron *Game) Next() {
	if tron.State.X < 5 {
		tron.State.X++
	} else {
		tron.Exited = true
	}
}

// UpdateWithInterval .
func (tron *Game) UpdateWithInterval(dur time.Duration, quit chan int) {
	ticker := time.NewTicker(dur)

	for range ticker.C {
		tron.Next()

		if tron.Exited {
			ticker.Stop()
			quit <- 0
		}
	}
}
