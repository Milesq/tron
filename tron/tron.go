package tron

import (
	"fmt"
	"time"

	"github.com/milesq/tron/tron/event"
	"github.com/milesq/tron/tron/utils"
)

// Game .
type Game struct {
	Exited           bool
	State            GameState
	PlayersDirection map[int]Vector
}

// NewGame .
func NewGame(cfg Config) Game {
	players := make(map[int]Trace)
	vectors := make(map[int]Vector)

	for _, player := range cfg.Players {
		players[player] = Trace{
			Point{
				utils.Random(0, cfg.Size[0]),
				utils.Random(0, cfg.Size[1]),
			},
		}

		var vec Vector
		vec.X = utils.ChooseRandom([]interface{}{-1, 1}).(int)
		vec.Y = utils.ChooseRandom([]interface{}{-1, 0, 1}).(int)
		vectors[player] = vec
	}

	return Game{
		false,
		GameState{
			MapSize:    cfg.Size,
			Players:    players,
			PlayerChar: cfg.PlayerChar,
			BorderChar: cfg.BorderChar,
		},
		vectors,
	}
}

// Emit .
func (tron *Game) Emit(ev event.Event, ctx int) {
	switch ev {
	case event.Exit:
		tron.Exited = true
	case event.Up:
		fmt.Println("up", ctx)
	case event.Right:
		fmt.Println("right", ctx)
	case event.Down:
		fmt.Println("down", ctx)
	case event.Left:
		fmt.Println("left", ctx)
	}
}

// Next .
func (tron *Game) Next() {
	for playerID := range tron.State.Players {
		trace := tron.State.Players[playerID]
		vector := tron.PlayersDirection[playerID]
		lastPos := trace[len(trace)-1]

		tron.State.Players[playerID] = append(trace, Point{lastPos.X + vector.X, lastPos.Y + vector.Y})
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
