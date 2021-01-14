package tron

import (
	"time"

	"github.com/milesq/tron/tron/event"
	"github.com/milesq/tron/tron/utils"
)

// Game .
type Game struct {
	Exited           bool
	State            GameState
	PlayersDirection map[string]Vector
}

// NewGame .
func NewGame(cfg Config) Game {
	players := make(map[string]Trace)

	for _, player := range cfg.Players {
		players[player] = Trace{
			Point{
				utils.Random(0, cfg.Size[0]),
				utils.Random(0, cfg.Size[1]),
			},
		}
	}

	return Game{
		State: GameState{
			MapSize: cfg.Size,
			Players: players,
		},
	}
}

// Emit .
func (tron *Game) Emit(ev event.Event) {
	switch ev {
	case event.Exit:
		tron.Exited = true
	}
}

// Next .
func (tron *Game) Next() {}

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
