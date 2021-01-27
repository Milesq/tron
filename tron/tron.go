package tron

import (
	"math/rand"
	"time"

	"github.com/milesq/tron/tron/event"
	"github.com/milesq/tron/tron/result"
)

// Game .
type Game struct {
	Exited           bool
	State            GameState
	PlayersDirection map[int]Vector
	Cfg              Config

	randInstance *rand.Rand
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

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	return Game{
		false,
		GameState{players, make(map[int]int)},
		vectors,
		cfg,
		r,
	}
}

// Emit .
func (tron *Game) Emit(ev event.Event, ctx int) {
	speed := tron.Cfg.PlayerSpeed
	v := Vector{speed, speed}

	switch ev {
	case event.Exit:
		tron.Exited = true
		return

	case event.Up:
		tron.PlayersDirection[ctx] = v.Mul(UP)
	case event.Right:
		tron.PlayersDirection[ctx] = v.Mul(RIGHT)
	case event.Down:
		tron.PlayersDirection[ctx] = v.Mul(DOWN)
	case event.Left:
		tron.PlayersDirection[ctx] = v.Mul(LEFT)
	}

	tron.randomAddPoints(3, ctx)
}

// UpdateWithInterval .
func (tron *Game) UpdateWithInterval(dur time.Duration, quit chan result.Result) {
	ticker := time.NewTicker(dur)

	for range ticker.C {
		tron.Next()

		if tron.Exited || len(tron.State.Players) == 1 {
			ticker.Stop()

			var won int

			for key := range tron.State.Points {
				won = key
				// there is only one player
				break
			}

			quit <- result.Result{
				Won:    won,
				Points: tron.State.Points,
			}
		}
	}
}

func (tron *Game) randomAddPoints(ev event.Event, ctx int) {
	_, playerExists := tron.State.Players[ctx]
	const chances = 6
	grantedPoints := randomWithChances([][2]int{{3, 2}, {4, 3}, {2, 4}, {1, 6}}, tron.randInstance)

	if playerExists && tron.randInstance.Intn(chances) == 1 {
		tron.State.Points[ctx] += grantedPoints
	}
}

func filled(size, init int) (list []int) {
	for i := 0; i < size; i++ {
		list = append(list, init)
	}

	return
}

func randomWithChances(list [][2]int, rd *rand.Rand) int {
	var finalList []int

	for _, item := range list {
		finalList = append(finalList, filled(item[0], item[1])...)
	}

	i := rd.Intn(len(finalList))

	return finalList[i]
}
