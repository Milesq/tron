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
	PlayersDirection map[int]Vector
	Cfg              Config
}

// Config .
type Config struct {
	Players     []int // []tm.Color
	PlayerSpeed float64
	Size        [2]int
	PlayerChar  byte
	BorderChar  byte
}

// NewGame .
func NewGame(cfg Config) Game {
	players := make(map[int]Trace)
	vectors := make(map[int]Vector)

	for _, player := range cfg.Players {
		players[player] = Trace{
			Point{
				float64(utils.Random(0, cfg.Size[0])),
				float64(utils.Random(0, cfg.Size[1])),
			},
		}

		var vec Vector
		vec.X = utils.ChooseRandom([]interface{}{-cfg.PlayerSpeed, cfg.PlayerSpeed}).(float64)
		vec.Y = utils.ChooseRandom([]interface{}{-cfg.PlayerSpeed, 0.0, cfg.PlayerSpeed}).(float64)
		vectors[player] = vec
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
