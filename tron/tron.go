package tron

import (
	"time"
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
