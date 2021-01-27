package main

import (
	"fmt"
	"time"

	"github.com/milesq/tron/displayer"
	"github.com/milesq/tron/tron/result"

	tm "github.com/buger/goterm"

	keyCtrl "github.com/milesq/tron/controllers/keyboard"
	"github.com/milesq/tron/tron"
)

func main() {
	tm.Clear()

	game := tron.NewGame(tron.Config{
		Players:     []int{tm.RED, tm.BLUE, tm.YELLOW},
		PlayerSpeed: 0.1,
		Size:        [2]int{100 | tm.PCT, 30},
		PlayerChar:  "•",
		BorderChar:  "█",
	})

	quit := make(chan result.Result)

	go game.UpdateWithInterval(time.Second/30, quit)

	go func() {
		displayInterval := time.NewTicker(time.Second / 40)

		for range displayInterval.C {
			tm.Clear()
			displayer.ConsoleDisplayer(game)
			tm.Flush()
		}
	}()

	go keyCtrl.Controller(&game)

	score := <-quit

	fmt.Println("score: ", score)
}
