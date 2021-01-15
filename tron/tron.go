package tron

import (
	"time"

	"github.com/milesq/tron/tron/event"
)

// Game .
type Game struct {
	Exited           bool
	State            GameState
	PlayersDirection map[int]Vector
	Cfg              Config
}

// Config .
type Config struct {
	Players     []int // []tm.Color
	PlayerSpeed float64
	Size        [2]int
	PlayerChar  string
	BorderChar  string
}

// NewGame .
func NewGame(cfg Config) Game {
	players := make(map[int]Trace)
	vectors := make(map[int]Vector)

	screenPart := float64(cfg.Size[1]/len(cfg.Players) + 1)
	const padding = 10

	for i, player := range cfg.Players {
		players[player] = Trace{
			Point{
				5,
				screenPart + float64(i)*padding,
			},
		}

		vectors[player] = Vector{cfg.PlayerSpeed, 0}
	}

	return Game{
		false,
		GameState{players},
		vectors,
		cfg,
	}
}

// Emit .
func (tron *Game) Emit(ev event.Event, ctx int) {
	speed := tron.Cfg.PlayerSpeed
	v := Vector{speed, speed}

	switch ev {
	case event.Exit:
		tron.Exited = true

	case event.Up:
		tron.PlayersDirection[ctx] = v.Mul(UP)
	case event.Right:
		tron.PlayersDirection[ctx] = v.Mul(RIGHT)
	case event.Down:
		tron.PlayersDirection[ctx] = v.Mul(DOWN)
	case event.Left:
		tron.PlayersDirection[ctx] = v.Mul(LEFT)
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
